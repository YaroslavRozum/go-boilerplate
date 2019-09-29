package models

var DefaultProductMapper *ProductMapper

func initProductMapper() {
	DefaultProductMapper = &ProductMapper{}
}

type Product struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Price int    `json:"price" db:"price"`
}

type ProductMapper struct{}

func (uM *ProductMapper) FindAll(where interface{}, limit, offset uint64, args ...interface{}) ([]*Product, error) {
	selectBuilder := QueryBuilder.
		Select("*").
		From("products").
		Where(where, args...)
	if limit != 0 {
		selectBuilder = selectBuilder.Limit(limit)
	}
	if offset != 0 {
		selectBuilder = selectBuilder.Offset(offset)
	}

	query, queryArgs, _ := selectBuilder.ToSql()

	rows, err := DB.Queryx(query, queryArgs...)
	if err != nil {
		return nil, err
	}

	products := initProductsByLimit(limit)
	for rows.Next() {
		product := &Product{}
		err := rows.StructScan(product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func initProductsByLimit(limit uint64) []*Product {
	if limit != 0 {
		return make([]*Product, 0, limit)
	}
	return []*Product{}
}
