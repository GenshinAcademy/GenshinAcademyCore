package models

type AcademyId uint32 // @name AcademyId

const UNDEFINED_ID AcademyId = 0

type AcademyModel struct {
	Id AcademyId `extensions:"x-order=0"`
}
