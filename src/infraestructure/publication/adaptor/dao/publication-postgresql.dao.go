package dao

import (
	"fmt"
	"golang-gingonic-hex-architecture/src/application/publication/query/dto"
	"golang-gingonic-hex-architecture/src/infraestructure/publication/entity"

	"gorm.io/gorm"
)

type DaoPublicationPostgreSql struct {
	daoPublication *gorm.DB
}

func NewDaoPublicationPostgreSql(conn *gorm.DB) *DaoPublicationPostgreSql {
	return &DaoPublicationPostgreSql{
		daoPublication: conn.Model(&entity.Publication{}),
	}
}

func (dap *DaoPublicationPostgreSql) List() []*dto.PublicationDto {
	var publications []*dto.PublicationDto
	dap.daoPublication.Raw("SELECT * FROM publications").Scan(&publications)
	return publications
}

func (dap *DaoPublicationPostgreSql) ListByFilters(filters dto.FilterPublicationsDto) []*dto.PublicationDto {
	var publications []*dto.PublicationDto
	fmt.Println("Filt", filters, len(publications))
	statement := dap.daoPublication
	if filters.Limit != 0 {
		statement.Limit(filters.Limit)
	}
	if filters.Offset != 0 {
		statement.Offset(filters.Offset)
	}
	for _, v := range filters.Type {
		fmt.Println("iteration", v.String())
		statement.Or("type = ?", v.String())
	}
	statement.Find(&publications)
	for _, v := range publications {
		fmt.Println("ar", &v.Type)
	}
	return publications
}
