package models

import sq "github.com/Masterminds/squirrel"

var QueryBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
