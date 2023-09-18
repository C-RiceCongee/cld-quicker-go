// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"cld-quicker-go/model/models"
)

func newFile(db *gorm.DB, opts ...gen.DOOption) file {
	_file := file{}

	_file.fileDo.UseDB(db, opts...)
	_file.fileDo.UseModel(&models.File{})

	tableName := _file.fileDo.TableName()
	_file.ALL = field.NewAsterisk(tableName)
	_file.ID = field.NewInt64(tableName, "id")
	_file.FileID = field.NewInt64(tableName, "file_id")
	_file.UserID = field.NewInt64(tableName, "user_id")
	_file.UploadType = field.NewString(tableName, "upload_type")
	_file.SceneType = field.NewString(tableName, "scene_type")
	_file.BaseURL = field.NewString(tableName, "base_url")
	_file.Path = field.NewString(tableName, "path")
	_file.Name = field.NewString(tableName, "name")
	_file.Extension = field.NewString(tableName, "extension")
	_file.Size = field.NewInt32(tableName, "size")
	_file.Year = field.NewInt32(tableName, "year")
	_file.Month = field.NewInt32(tableName, "month")
	_file.Day = field.NewInt32(tableName, "day")
	_file.UploadIP = field.NewString(tableName, "upload_ip")
	_file.Status = field.NewInt32(tableName, "status")
	_file.CreateTime = field.NewTime(tableName, "create_time")
	_file.UpdateTime = field.NewTime(tableName, "update_time")

	_file.fillFieldMap()

	return _file
}

type file struct {
	fileDo

	ALL        field.Asterisk
	ID         field.Int64
	FileID     field.Int64
	UserID     field.Int64
	UploadType field.String // 上传类型[七牛云，本地..]
	SceneType  field.String // 场景类别[风景，人物...]
	BaseURL    field.String // url(本地路径就是/，非本地就存远程URL)
	Path       field.String // 本地路径 /home/xxx 或者是远程的域名
	Name       field.String // 文件原始名
	Extension  field.String // 扩展名
	Size       field.Int32  // 长度
	Year       field.Int32  // 年份
	Month      field.Int32  // 月份
	Day        field.Int32  // 日
	UploadIP   field.String // 上传者ip
	Status     field.Int32  // 状态[-1:删除;0:禁用;1启用]
	CreateTime field.Time   // 数据创建时间
	UpdateTime field.Time   // 数据更新时间

	fieldMap map[string]field.Expr
}

func (f file) Table(newTableName string) *file {
	f.fileDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f file) As(alias string) *file {
	f.fileDo.DO = *(f.fileDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *file) updateTableName(table string) *file {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.FileID = field.NewInt64(table, "file_id")
	f.UserID = field.NewInt64(table, "user_id")
	f.UploadType = field.NewString(table, "upload_type")
	f.SceneType = field.NewString(table, "scene_type")
	f.BaseURL = field.NewString(table, "base_url")
	f.Path = field.NewString(table, "path")
	f.Name = field.NewString(table, "name")
	f.Extension = field.NewString(table, "extension")
	f.Size = field.NewInt32(table, "size")
	f.Year = field.NewInt32(table, "year")
	f.Month = field.NewInt32(table, "month")
	f.Day = field.NewInt32(table, "day")
	f.UploadIP = field.NewString(table, "upload_ip")
	f.Status = field.NewInt32(table, "status")
	f.CreateTime = field.NewTime(table, "create_time")
	f.UpdateTime = field.NewTime(table, "update_time")

	f.fillFieldMap()

	return f
}

func (f *file) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *file) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 17)
	f.fieldMap["id"] = f.ID
	f.fieldMap["file_id"] = f.FileID
	f.fieldMap["user_id"] = f.UserID
	f.fieldMap["upload_type"] = f.UploadType
	f.fieldMap["scene_type"] = f.SceneType
	f.fieldMap["base_url"] = f.BaseURL
	f.fieldMap["path"] = f.Path
	f.fieldMap["name"] = f.Name
	f.fieldMap["extension"] = f.Extension
	f.fieldMap["size"] = f.Size
	f.fieldMap["year"] = f.Year
	f.fieldMap["month"] = f.Month
	f.fieldMap["day"] = f.Day
	f.fieldMap["upload_ip"] = f.UploadIP
	f.fieldMap["status"] = f.Status
	f.fieldMap["create_time"] = f.CreateTime
	f.fieldMap["update_time"] = f.UpdateTime
}

func (f file) clone(db *gorm.DB) file {
	f.fileDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f file) replaceDB(db *gorm.DB) file {
	f.fileDo.ReplaceDB(db)
	return f
}

type fileDo struct{ gen.DO }

func (f fileDo) Debug() *fileDo {
	return f.withDO(f.DO.Debug())
}

func (f fileDo) WithContext(ctx context.Context) *fileDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fileDo) ReadDB() *fileDo {
	return f.Clauses(dbresolver.Read)
}

func (f fileDo) WriteDB() *fileDo {
	return f.Clauses(dbresolver.Write)
}

func (f fileDo) Session(config *gorm.Session) *fileDo {
	return f.withDO(f.DO.Session(config))
}

func (f fileDo) Clauses(conds ...clause.Expression) *fileDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fileDo) Returning(value interface{}, columns ...string) *fileDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fileDo) Not(conds ...gen.Condition) *fileDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fileDo) Or(conds ...gen.Condition) *fileDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fileDo) Select(conds ...field.Expr) *fileDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fileDo) Where(conds ...gen.Condition) *fileDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fileDo) Order(conds ...field.Expr) *fileDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fileDo) Distinct(cols ...field.Expr) *fileDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fileDo) Omit(cols ...field.Expr) *fileDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fileDo) Join(table schema.Tabler, on ...field.Expr) *fileDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fileDo) LeftJoin(table schema.Tabler, on ...field.Expr) *fileDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fileDo) RightJoin(table schema.Tabler, on ...field.Expr) *fileDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fileDo) Group(cols ...field.Expr) *fileDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fileDo) Having(conds ...gen.Condition) *fileDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fileDo) Limit(limit int) *fileDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fileDo) Offset(offset int) *fileDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fileDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *fileDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fileDo) Unscoped() *fileDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fileDo) Create(values ...*models.File) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fileDo) CreateInBatches(values []*models.File, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fileDo) Save(values ...*models.File) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fileDo) First() (*models.File, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.File), nil
	}
}

func (f fileDo) Take() (*models.File, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.File), nil
	}
}

func (f fileDo) Last() (*models.File, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.File), nil
	}
}

func (f fileDo) Find() ([]*models.File, error) {
	result, err := f.DO.Find()
	return result.([]*models.File), err
}

func (f fileDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.File, err error) {
	buf := make([]*models.File, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fileDo) FindInBatches(result *[]*models.File, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fileDo) Attrs(attrs ...field.AssignExpr) *fileDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fileDo) Assign(attrs ...field.AssignExpr) *fileDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fileDo) Joins(fields ...field.RelationField) *fileDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fileDo) Preload(fields ...field.RelationField) *fileDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fileDo) FirstOrInit() (*models.File, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.File), nil
	}
}

func (f fileDo) FirstOrCreate() (*models.File, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.File), nil
	}
}

func (f fileDo) FindByPage(offset int, limit int) (result []*models.File, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fileDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fileDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fileDo) Delete(models ...*models.File) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fileDo) withDO(do gen.Dao) *fileDo {
	f.DO = *do.(*gen.DO)
	return f
}
