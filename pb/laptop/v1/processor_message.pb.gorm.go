package pb

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm "github.com/jinzhu/gorm"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	strings "strings"
)

type CpuORM struct {
	Brand         string
	Id            uint32 `gorm:"primary_key;auto_increment"`
	LaptopId      *string
	MaxGhz        float64
	MinGhz        float64
	Name          string
	NumberCores   uint32
	NumberThreads uint32
}

// TableName overrides the default tablename generated by GORM
func (CpuORM) TableName() string {
	return "cpus"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Cpu) ToORM(ctx context.Context) (CpuORM, error) {
	to := CpuORM{}
	var err error
	if prehook, ok := interface{}(m).(CpuWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Brand = m.Brand
	to.Name = m.Name
	to.NumberCores = m.NumberCores
	to.NumberThreads = m.NumberThreads
	to.MinGhz = m.MinGhz
	to.MaxGhz = m.MaxGhz
	if posthook, ok := interface{}(m).(CpuWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *CpuORM) ToPB(ctx context.Context) (Cpu, error) {
	to := Cpu{}
	var err error
	if prehook, ok := interface{}(m).(CpuWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Brand = m.Brand
	to.Name = m.Name
	to.NumberCores = m.NumberCores
	to.NumberThreads = m.NumberThreads
	to.MinGhz = m.MinGhz
	to.MaxGhz = m.MaxGhz
	if posthook, ok := interface{}(m).(CpuWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Cpu the arg will be the target, the caller the one being converted from

// CpuBeforeToORM called before default ToORM code
type CpuWithBeforeToORM interface {
	BeforeToORM(context.Context, *CpuORM) error
}

// CpuAfterToORM called after default ToORM code
type CpuWithAfterToORM interface {
	AfterToORM(context.Context, *CpuORM) error
}

// CpuBeforeToPB called before default ToPB code
type CpuWithBeforeToPB interface {
	BeforeToPB(context.Context, *Cpu) error
}

// CpuAfterToPB called after default ToPB code
type CpuWithAfterToPB interface {
	AfterToPB(context.Context, *Cpu) error
}

type GpuORM struct {
	Brand    string
	LaptopId *string
	MaxGhz   float64
	MinGhz   float64
	Name     string
}

// TableName overrides the default tablename generated by GORM
func (GpuORM) TableName() string {
	return "gpus"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Gpu) ToORM(ctx context.Context) (GpuORM, error) {
	to := GpuORM{}
	var err error
	if prehook, ok := interface{}(m).(GpuWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Brand = m.Brand
	to.Name = m.Name
	to.MinGhz = m.MinGhz
	to.MaxGhz = m.MaxGhz
	if posthook, ok := interface{}(m).(GpuWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *GpuORM) ToPB(ctx context.Context) (Gpu, error) {
	to := Gpu{}
	var err error
	if prehook, ok := interface{}(m).(GpuWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Brand = m.Brand
	to.Name = m.Name
	to.MinGhz = m.MinGhz
	to.MaxGhz = m.MaxGhz
	if posthook, ok := interface{}(m).(GpuWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Gpu the arg will be the target, the caller the one being converted from

// GpuBeforeToORM called before default ToORM code
type GpuWithBeforeToORM interface {
	BeforeToORM(context.Context, *GpuORM) error
}

// GpuAfterToORM called after default ToORM code
type GpuWithAfterToORM interface {
	AfterToORM(context.Context, *GpuORM) error
}

// GpuBeforeToPB called before default ToPB code
type GpuWithBeforeToPB interface {
	BeforeToPB(context.Context, *Gpu) error
}

// GpuAfterToPB called after default ToPB code
type GpuWithAfterToPB interface {
	AfterToPB(context.Context, *Gpu) error
}

// DefaultCreateCpu executes a basic gorm create call
func DefaultCreateCpu(ctx context.Context, in *Cpu, db *gorm.DB) (*Cpu, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type CpuORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CpuORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadCpu(ctx context.Context, in *Cpu, db *gorm.DB) (*Cpu, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm1.ApplyFieldSelection(ctx, db, nil, &CpuORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := CpuORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(CpuORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type CpuORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CpuORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CpuORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteCpu(ctx context.Context, in *Cpu, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&CpuORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type CpuORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CpuORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteCpuSet(ctx context.Context, in []*Cpu, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []uint32{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&CpuORM{})).(CpuORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&CpuORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&CpuORM{})).(CpuORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type CpuORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Cpu, *gorm.DB) (*gorm.DB, error)
}
type CpuORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Cpu, *gorm.DB) error
}

// DefaultStrictUpdateCpu clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateCpu(ctx context.Context, in *Cpu, db *gorm.DB) (*Cpu, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateCpu")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &CpuORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(CpuORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type CpuORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CpuORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CpuORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchCpu executes a basic gorm update call with patch behavior
func DefaultPatchCpu(ctx context.Context, in *Cpu, updateMask *field_mask.FieldMask, db *gorm.DB) (*Cpu, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj Cpu
	var err error
	if hook, ok := interface{}(&pbObj).(CpuWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadCpu(ctx, &Cpu{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(CpuWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskCpu(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(CpuWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateCpu(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(CpuWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type CpuWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Cpu, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CpuWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Cpu, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CpuWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Cpu, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type CpuWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Cpu, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetCpu executes a bulk gorm update call with patch behavior
func DefaultPatchSetCpu(ctx context.Context, objects []*Cpu, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*Cpu, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Cpu, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchCpu(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskCpu patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskCpu(ctx context.Context, patchee *Cpu, patcher *Cpu, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Cpu, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Brand" {
			patchee.Brand = patcher.Brand
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"NumberCores" {
			patchee.NumberCores = patcher.NumberCores
			continue
		}
		if f == prefix+"NumberThreads" {
			patchee.NumberThreads = patcher.NumberThreads
			continue
		}
		if f == prefix+"MinGhz" {
			patchee.MinGhz = patcher.MinGhz
			continue
		}
		if f == prefix+"MaxGhz" {
			patchee.MaxGhz = patcher.MaxGhz
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListCpu executes a gorm list call
func DefaultListCpu(ctx context.Context, db *gorm.DB) ([]*Cpu, error) {
	in := Cpu{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &CpuORM{}, &Cpu{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []CpuORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(CpuORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Cpu{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type CpuORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CpuORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type CpuORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]CpuORM) error
}

// DefaultCreateGpu executes a basic gorm create call
func DefaultCreateGpu(ctx context.Context, in *Gpu, db *gorm.DB) (*Gpu, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GpuORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GpuORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type GpuORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type GpuORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

// DefaultApplyFieldMaskGpu patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskGpu(ctx context.Context, patchee *Gpu, patcher *Gpu, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*Gpu, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedMemory bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Brand" {
			patchee.Brand = patcher.Brand
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"MinGhz" {
			patchee.MinGhz = patcher.MinGhz
			continue
		}
		if f == prefix+"MaxGhz" {
			patchee.MaxGhz = patcher.MaxGhz
			continue
		}
		if !updatedMemory && strings.HasPrefix(f, prefix+"Memory.") {
			if patcher.Memory == nil {
				patchee.Memory = nil
				continue
			}
			if patchee.Memory == nil {
				patchee.Memory = &Memory{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"Memory."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.Memory, patchee.Memory, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"Memory" {
			updatedMemory = true
			patchee.Memory = patcher.Memory
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListGpu executes a gorm list call
func DefaultListGpu(ctx context.Context, db *gorm.DB) ([]*Gpu, error) {
	in := Gpu{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GpuORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm1.ApplyCollectionOperators(ctx, db, &GpuORM{}, &Gpu{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GpuORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	ormResponse := []GpuORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GpuORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Gpu{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type GpuORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type GpuORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type GpuORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]GpuORM) error
}