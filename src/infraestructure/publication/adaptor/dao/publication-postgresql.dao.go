package dao

import (
	commentDto "golang-gingonic-hex-architecture/src/application/comment/query/dto"
	"golang-gingonic-hex-architecture/src/application/publication/query/dto"
	ce "golang-gingonic-hex-architecture/src/infraestructure/comment/entity"
	"golang-gingonic-hex-architecture/src/infraestructure/publication/entity"
	"strconv"

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

func (dap DaoPublicationPostgreSql) GetOneById(id int) *dto.PublicationDto {
	parsedId := strconv.Itoa(id)
	var rawSql = "SELECT " +
		"publications.id, " +
		"publications.title, " +
		"publications.content, " +
		"publications.description, " +
		"publications.type, " +
		"publications.categories, " +
		"publications.updated_at, " +
		"publications.witer_user_id, " +
		"comments.id, " +
		"comments.comment_title, " +
		"comments.comment_content " +
		"FROM publications " +
		"LEFT JOIN comments ON publications.id = comments.publication_id " +
		"WHERE publications.deleted_at IS NULL AND publications.id = " + parsedId
	rows, _ := dap.daoPublication.Raw(rawSql).Rows()
	var pp *dto.PublicationDto
	var comments []commentDto.CommentDto = make([]commentDto.CommentDto, 0)
	p := &entity.Publication{}
	for rows.Next() {
		comment := ce.Comment{}
		rows.Scan(
			&p.Id,
			&p.Title,
			&p.Content,
			&p.Description,
			&p.Type,
			&p.Categories,
			&p.WrittenAt,
			&p.WiterUserId,
			&comment.Id,
			&comment.CommentTitle,
			&comment.CommentContent,
		)
		c := commentDto.CommentDto{
			CommentTitle:   comment.CommentTitle,
			CommentContent: comment.CommentContent,
			Id:             comment.Id,
		}
		comments = append(comments, c)
	}

	pp = &dto.PublicationDto{
		Id:          p.Id,
		Title:       p.Title,
		Description: p.Description,
		Content:     p.Content,
		Type:        p.Type,
		Categories:  p.Categories,
		WrittenAt:   p.WrittenAt,
		WiterUserId: p.WiterUserId,
		Comments:    comments,
	}
	return pp
}

func (dap DaoPublicationPostgreSql) List() []*dto.PublicationDto {
	var publications []*dto.PublicationDto

	var rawSQL = "SELECT * FROM publications p WHERE p.deleted_at IS NULL"
	dap.daoPublication.Raw(rawSQL).Find(&publications)

	return publications
}

func (dap DaoPublicationPostgreSql) ListByFilters(filters dto.FilterPublicationsDto) []*dto.PublicationDto {
	var publications []*dto.PublicationDto
	sqlRaw := "SELECT * FROM publications WHERE ("

	length := len(filters.Type)
	for index, v := range filters.Type {
		if index != length-1 {
			sqlRaw += "type = '" + v.String() + "' OR "
		} else {
			sqlRaw += "type = '" + v.String() + "' )"
		}
	}
	dap.daoPublication.Raw(sqlRaw).Scan(&publications)
	return publications
}

func (dap DaoPublicationPostgreSql) Search(text string) []*dto.PublicationDto {
	var publications []*dto.PublicationDto

	var rawSQL = "SELECT * FROM publications p, unnest(categories) c WHERE "
	rawSQL += "(c LIKE ? OR p.title LIKE ? OR p.description LIKE ? OR p.content LIKE ? )"
	dap.daoPublication.Raw(
		rawSQL,
		"%"+text+"%",
		"%"+text+"%",
		"%"+text+"%",
		"%"+text+"%",
	).Scan(&publications)

	return publications
}
