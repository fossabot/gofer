# Grid

The grid package converts any `grid.Source` into a CRUD backend.

* All primary-, foreign, references and polymorphic fields are removed by default
* Relation will be displayed depth max 1 at the moment.
* `belongsTo` and `manyToMany` will be select boxes on the frontend.
* Errors will be set as controller errors.
* Field validation will happen automatically on front- and backend.
* Frontend fields will be rendered automatically by field type.
* Developer friendly. Every source which implements the `grid.Source` can be used.

## Usage

```go 
func (c MyController) User(){
    g := grid.New(c,grid.Orm(model),nil)
    //...
    g.Render()
} 
```

## New

Creates a new grid instance. The first argument must be the `controller`, the second is the `grid.Source` and the third
is the `grid.config` which is optional.

The grid will be cached, to avoid re-creating the grid fields. The cache key will be the configured grid ID.

```go
g := grid.New(c, grid.Orm(model), nil)
```

## Config

The grid can be fully configured. If the configuration should be changed after init dynamically, the `Scope.Config()`
can be used.

| Name        | Default |Description                                                       | 
|-------------|----|-------------------------------------------------------------------|
| ID          | `controller:action`    | Unique name for the grid. This is used as cache key.              |
| Title       | `{ID}-title`   |  Title of the grid.                 |
| Description | `{ID}-description`   | Description of the grid.              |
| Policy      | `orm.WHITELIST`   | If the Policy is `WHITELIST`, the fields have to be set explicit. |
| Exports     | `csv`, `pdf`   | Slice of names of render types.                                      |
| Action      |    | see ACTION                                                        |
| Filter      |    | see FILTER                                                        |
| Translation      |`false`     | If set to true, the title will be the orm namespace-title. same goes for the description.                                                    |

**Action**

| Name        |  Default |  Description                                                       |
|-------------|-----|----------------------------------------------------------------|
| Right          | `true`   | Defines where the action (details,edit,delete) column should be displayed on the grid table.            |
| AllowDetails       | `false` |  Allow details mode.                     |
| AllowCreate | `true`  |  Allow create mode.                 |
| AllowUpdate      | `true`  |  Allow update mode. |
| AllowDelete      | `true`  |  Allow delete mode.                                                       |

**Filter**

| Name        |  Default |  Description                                                       |
|-------------|-----|----------------------------------------------------------------|
| Allow          | `true`   | Allow filter.            |
| ShowQuickFilter       | `true` |  Show the quick filter bar.                        |
| ShowCustomFilter | `true`  |  Show the custom filter button/view.                  |
| AllowedRowsPerPage      | `-1`,`5`,`10`,`15`,`25`,`50`  |  The allowed rows per page. |
| DefaultRowsPerPage      | `15`  |  Default rows per page.                                                     |

## Mode

The grid mode is defined by the `http.Method` and `controller.Params`.

| Mode        |  http.Method |  Param                                                       |
|-------------|-----|----------------------------------------------------------------|
| `grid.FeTable`          | `GET`   |             |
| `grid.FeFilter`          | `GET`   | `mode=filter`            |
| `grid.SrcCallback`          | `GET`   | `mode=callback`            |
| `grid.FeDetails`          | `GET`   | `mode=details`            |
| `grid.FeCreate`          | `GET`   | `mode=create`            |
| `grid.FeUpdate`          | `GET`   | `mode=update`            |
| `grid.FeExport`          | `GET`   | `mode=export`            |
| `grid.SrcCreate`          | `POST`   |           |
| `grid.SrcUpdate`          | `PUT`   |           |
| `grid.SrcDelete`          | `DELETE`   |           |

## Field

Will return the grid field by name. If the field does not exist, an empty field with an error will return.

```go
field := grid.Field("ID")
```

A field can be configured by the following functions. Each function returns itself, this way it can be chained. If an
error occures, the fields error will be set. Error will be handled in `grid.Render`.

The configuration for `SetPosition`, `SetTitle`, `SetDescription`, `SetRemove`, `SetHidden` and `SetView` must be set
with `grid.NewValue()`. Like this, values can be defined for different grid modes.

```go
grid.Field("ID").SetTitle(grid.NewValue("ID").SetDetails("Identifier"))
// grid mode: table, update, create will have the title "ID" 
// and details will have the title "Identifier".
```

| Function        |  available frontend |  Description | 
|-------------|-----|-----|
| `Name`, `SetName`   | x | Will set the fields name. The name is used in the frontend as id.   |
| `Primary`, `SetPrimary`  | x | Will define if the field is a primary key.  |
| `Type`, `SetType`   | x | Defines the field type.  |
| `Title`, `SetTitle`   |x | Will set the fields title.  |
| `Description`, `SetDescription`  |x | Will set the fields description.  |
| `Position`, `SetPosition`  |x | Will set the fields position.  |
| `Removed`, `SetRemove`  |x | Will flag the field as removed.  |
| `Hidden`, `SetHidden`  | x| Will set the field as hidden.  |
| `View`, `SetView`   |x| Will set a custom frontend view component for the field.  |
| `ReadOnly`, `SetReadOnly`  | x| Will set the field as read only.  |
| `Sort`, `SetSort`   |x| Will allow the sorting of the field and set the condition field name. |
| `Filter`, `SetFilter`   |x| Will allow the filtering of the field and set the condition operator and field name. |
| `GroupAble`, `SetGroupAble`   |x| Will set the field as group able.  |
| `Options`, `Option`, `SetOption`  | x| Will add a option for the field.  |
| `Relation`, `SetRelation`   |x| Will define the field as relation  |
| `Field`   || Will return a field by name. If the field was not found, an field error will be set. (relation)  |
| `Fields`, `SetFields`   |x| Will return all child fields. (relation)  |
| `Error`  || Will return the field error. |

**Field types**

| Name        |  implemented in frontend |  Description |
|-------------|-----|-----|
| `Bool`   |  | Checkbox   |
| `Integer`   |  | Input-Integer   |
| `Float`   |  | Input-Numeric   |
| `Text`   |  | Input-Text   |
| `TextArea`   |  | TextArea   |
| `Time`   |  | Input   |
| `Date`   |  | Datepicker   |
| `DateTime`   |  | Datepicker+Input   |
| `Select`   |  | Select   |
| `MultiSelect`   |  | Select   |
| `belongsTo`   |  | Select   |
| `hasOne`   |  | Inputs   |
| `hasMany`   |  | Inputs   |
| `m2m`   |  | Select   |

**Options**

| Name        |  value |  Description |
|-------------|-----|-----|
| `DecoratorOption`   | `string`,`boolean` | a field name can be used {{Name}}. If the second parameter is set to true, HTML will not be escaped in the frontend.   |

**Callbacks**

| Name        |  value |  Description |
|-------------|-----|-----|
| `Select`   | `?` |   |

TODO: Validate

## Scope

The scope will return some helper functions.

### Source

Will return the grid source.

```go
src := scrope.Source()
```

### Config

Will return a pointer to the grid config. For dynamically configuration of the grid.

```go
cfg := scrope.Config()
```

### Fields

Will return all configured grid fields.

```go
fields := scrope.Fields()
```

### PrimaryFields

Will return all defined primary fields of the grid.

```go
primaryFields := scrope.PrimaryFields()
```

### Controller

Will return the grid controller instance.

```go
ctrl := scrope.Controller()
```

## Render

Will render the grid by the actual grid mode.

| Mode        | set in frontend data |  Description | 
|-------------|-----|-----|
| `grid.SrcCallback`   | `data`  | The source callback function is called. as first param the requested callback will be set as string. |
| `grid.SrcCreate`   |   | The source create function is called. | 
| `grid.SrcUpdate`     || The source update function is called. | 
| `grid.SrcDelete`     || The condition first will be called to ensure the correct primary key. The source delete function is called.| 
| `grid.FeTable`    | `pagination`, `head`, `data`, `config`| ConditionAll is called to create the condition. Add header/pagination if its not excluded by param. The source all function is called. Add config and result to the controller. call the defined render type.| 
| `grid.FeExport`     | `head`, `data`, `config`| Same as FeTable but without the pagination and limit.|
| `grid.FeCreate`    |`head` | add header data. | 
| `grid.FeDetails`,`grid.FeUpdate`    | `head`, `data`| add header data. call conditionFirst. fetch the entry by the given id and set the controller data. | 
| `grid.FeFilter`    | | TODO | 

## Orm

With the orm function an `orm.Interface` will be converted to a `grid.Source` and can be used out of the box.

```go
g := grid.New(ctrl, grid.Orm(model), nil)
```

## Source interface

To create your own source, you have to implement the `grid.Source`.

```go
type Source interface {
Cache() cache.Manager

PreInit(Grid) error
Init(Grid) error
Fields(Grid) ([]Field, error)
UpdatedFields(Grid) error

Callback(string, Grid) (interface{}, error)
First(condition.Condition, Grid) (interface{}, error)
All(condition.Condition, Grid) (interface{}, error)
Create(Grid) (interface{}, error)
Update(Grid) error
Delete(condition.Condition, Grid) error
Count(condition.Condition, Grid) (int, error)

//Interface() interface{}
}
```
