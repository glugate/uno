package query

import (
	"testing"
)

func TestGetErr(t *testing.T) {
	sb := NewBuilder()

	_, err := sb.GetQuery()
	if err != ErrTableEmpty {
		t.Error("check err")
	}

	sb = NewBuilder()

	_, err = sb.GetInsert()
	if err != ErrTableEmpty {
		t.Error("check err")
	}

	sb = NewBuilder()

	_, err = sb.GetUpdate()
	if err != ErrTableEmpty {
		t.Error("check err")
	}

	sb = NewBuilder()

	_, err = sb.GetDelete()
	if err != ErrTableEmpty {
		t.Error("check err")
	}

	sb = NewBuilder()

	_, err = sb.Table("test").GetInsert()
	if err != ErrInsertEmpty {
		t.Error("check err")
	}

	sb = NewBuilder()

	_, err = sb.Table("test").GetUpdate()
	if err != ErrUpdateEmpty {
		t.Error("check err")
	}
}
func TestBuilderSelect(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT `name`,`age`,`email` FROM `test`"
	if sql != expect {
		t.Error("sql gen err")
	}
}

func TestBuilderSelectAll(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT * FROM `test`"
	if sql != expect {
		t.Error("sql gen err")
	}
}

func TestBuilderSelect2(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("count(`age`), username").
		GetQuery()
	if err != nil {
		t.Error(err)
	}
	expect := "SELECT count(`age`), username FROM `test`"
	if sql != expect {
		t.Error("sql gen err")
	}
}

func TestBuilderWhere(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		Where("`name`", "=", "jack").
		Where("`age`", ">=", 18).
		OrWhere("`name`", "like", "%admin%").
		GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT `name`,`age`,`email` FROM `test` WHERE `name` = ? AND `age` >= ? OR `name` like ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "jack" ||
		params[1].(int) != 18 ||
		params[2].(string) != "%admin%" {
		t.Error("params gen err")
	}
}

func TestBuilderWhereRaw(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		WhereRaw("`title` = ?", "hello").
		Where("`name`", "=", "jack").
		OrWhereRaw("(`age` = ? OR `age` = ?)", 22, 25).
		GetQuery()
	if err != nil {
		t.Error(err)
	}
	expect := "SELECT `name`,`age`,`email` FROM `test` WHERE `title` = ? AND `name` = ? OR (`age` = ? OR `age` = ?)"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "hello" {
		t.Error("params gen err")
	}
	if params[1].(string) != "jack" {
		t.Error("params gen err")
	}
	if params[2].(int) != 22 {
		t.Error("params gen err")
	}
	if params[3].(int) != 25 {
		t.Error("params gen err")
	}
}

func TestBuilderWhereIn(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		WhereIn("`id`", 1, 2, 3).
		OrWhereNotIn("`uid`", 2, 4).
		GetQuery()
	if err != nil {
		t.Error(err)
	}
	expect := "SELECT `name`,`age`,`email` FROM `test` WHERE `id` IN (?,?,?) OR `uid` NOT IN (?,?)"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 1 {
		t.Error("params gen err")
	}
	if params[1].(int) != 2 {
		t.Error("params gen err")
	}
	if params[2].(int) != 3 {
		t.Error("params gen err")
	}
	if params[3].(int) != 2 {
		t.Error("params gen err")
	}
	if params[4].(int) != 4 {
		t.Error("params gen err")
	}
}

func TestBuilderOrWhereIn(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		OrWhereIn("`id`", 1, 2, 3).
		WhereNotIn("`uid`", 2, 4).
		GetQuery()
	if err != nil {
		t.Error(err)
	}
	expect := "SELECT `name`,`age`,`email` FROM `test` WHERE `id` IN (?,?,?) AND `uid` NOT IN (?,?)"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 1 {
		t.Error("params gen err")
	}
	if params[1].(int) != 2 {
		t.Error("params gen err")
	}
	if params[2].(int) != 3 {
		t.Error("params gen err")
	}
	if params[3].(int) != 2 {
		t.Error("params gen err")
	}
	if params[4].(int) != 4 {
		t.Error("params gen err")
	}
}

func TestBuilderGroupBy(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		GroupBy("`email`", "`class`").
		GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT `name`,`age`,`email` FROM `test` GROUP BY `email`,`class`"
	if sql != expect {
		t.Error("sql gen err")
	}
}

func TestBuilderHaving(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		GroupBy("`email`", "`class`").
		Having("COUNT(`name`)", ">", 10).
		GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT `name`,`age`,`email` FROM `test` GROUP BY `email`,`class` HAVING COUNT(`name`) > ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 10 {
		t.Error("params gen err")
	}
}

func TestBuilderOrHaving(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		GroupBy("`email`", "`class`").
		Having("`name`", "=", "a").
		OrHaving("`age`", "=", 12).
		GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT `name`,`age`,`email` FROM `test` GROUP BY `email`,`class` HAVING `name` = ? OR `age` = ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "a" {
		t.Error("params gen err")
	}

	if params[1].(int) != 12 {
		t.Error("params gen err")
	}
}

func TestBuilderHavingNotGen(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		Having("`name`", "=", "a").
		GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT `name`,`age`,`email` FROM `test`"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if len(params) != 0 {
		t.Error("params gen err")
	}
}

func TestBuilderHavingRaw(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		GroupBy("`email`", "`class`").
		Having("`name`", "=", "a").
		HavingRaw("count(`email`) <= ?", 22).
		GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT `name`,`age`,`email` FROM `test` GROUP BY `email`,`class` HAVING `name` = ? AND count(`email`) <= ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "a" {
		t.Error("params gen err")
	}
	if params[1].(int) != 22 {
		t.Error("params gen err")
	}
}

func TestBuilderOrHavingRaw(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		GroupBy("`email`", "`class`").
		OrHavingRaw("count(`email`) <= ?", 22).
		GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT `name`,`age`,`email` FROM `test` GROUP BY `email`,`class` HAVING count(`email`) <= ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 22 {
		t.Error("params gen err")
	}
}
func TestBuilderOrderBy(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		OrderBy("ASC", "`age`").
		GetQuery()
	if err != nil {
		t.Error(err)
	}
	expect := "SELECT `name`,`age`,`email` FROM `test` ORDER BY `age` ASC"
	if sql != expect {
		t.Error("sql gen err")
	}

}

func TestBuilderOrderBy2(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		OrderBy("ASC", "`age`", "`class`").
		GetQuery()
	if err != nil {
		t.Error(err)
	}
	expect := "SELECT `name`,`age`,`email` FROM `test` ORDER BY `age`,`class` ASC"
	if sql != expect {
		t.Error("sql gen err")
	}

}

func TestBuilderLimit(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		Limit(1, 10).
		GetQuery()
	if err != nil {
		t.Error(err)
	}
	expect := "SELECT `name`,`age`,`email` FROM `test` LIMIT ? OFFSET ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()
	if params[0].(int) != 10 {
		t.Error("params gen err")
	}
	if params[1].(int) != 1 {
		t.Error("params gen err")
	}
}

func TestBuilderQuery(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`name`", "`age`", "`email`").
		Where("`name`", "=", "jack").
		Where("`age`", ">=", 18).
		OrderBy("DESC", "`age`").
		Limit(1, 10).
		GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT `name`,`age`,`email` FROM `test` WHERE `name` = ? AND `age` >= ? ORDER BY `age` DESC LIMIT ? OFFSET ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(string) != "jack" ||
		params[1].(int) != 18 ||
		params[2].(int) != 10 ||
		params[3].(int) != 1 {
		t.Error("params gen err")
	}
}

func TestJoin(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`test`.`name`, `test`.`age`, `test2`.`status`").
		JoinRaw("LEFT JOIN `test2` ON `test`.`class` = `test2`.`class`").
		Where("`test`.`age`", ">=", 18).
		GetQuery()
	if err != nil {
		t.Error(err)
	}
	expect := "SELECT `test`.`name`, `test`.`age`, `test2`.`status` FROM `test`" +
		" LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` WHERE `test`.`age` >= ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 18 {
		t.Error("params gen err")
	}
}

func TestJoinWithParams(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Select("`test`.`name`, `test`.`age`, `test2`.`status`").
		JoinRaw("LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` AND `test`.`num` = ?", 2333).
		Where("`test`.`age`", ">=", 18).
		GetQuery()
	if err != nil {
		t.Error(err)
	}
	expect := "SELECT `test`.`name`, `test`.`age`, `test2`.`status` FROM `test`" +
		" LEFT JOIN `test2` ON `test`.`class` = `test2`.`class` AND `test`.`num` = ? WHERE `test`.`age` >= ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 2333 || params[1].(int) != 18 {
		t.Error("params gen err")
	}
}

func TestJoin2(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test` as t1").
		Select("`t1`.`name`", "`t1`.`age`", "`t2`.`status`", "`t3`.`address`").
		JoinRaw("LEFT JOIN `test2` as `t2` ON `t1`.`class` = `t2`.`class`").
		JoinRaw("INNER JOIN `test3` as t3 ON `t1`.`email` = `t3`.`email`").
		Where("`t1`.`age`", ">=", 18).
		GetQuery()
	if err != nil {
		t.Error(err)
	}
	expect := "SELECT `t1`.`name`,`t1`.`age`,`t2`.`status`,`t3`.`address` FROM `test` as t1" +
		" LEFT JOIN `test2` as `t2` ON `t1`.`class` = `t2`.`class`" +
		" INNER JOIN `test3` as t3 ON `t1`.`email` = `t3`.`email` WHERE `t1`.`age` >= ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 18 {
		t.Error("params gen err")
	}
}

func TestJoin3(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test` as t1").
		Select("`t1`.`name`", "`t1`.`age`", "`t2`.`status`", "`t3`.`address`").
		JoinRaw("LEFT JOIN `test2` as `t2` ON `t1`.`class` = `t2`.`class`").
		JoinRaw("INNER JOIN `test3` as t3 ON `t1`.`email` = `t3`.`email`").
		Where("`t1`.`age`", ">=", 18).
		GroupBy("`t1`.`age`").
		Having("COUNT(`t1`.`age`)", ">", 2).
		OrderBy("DESC", "`t1`.`age`").
		Limit(1, 10).
		GetQuery()
	if err != nil {
		t.Error(err)
	}

	expect := "SELECT `t1`.`name`,`t1`.`age`,`t2`.`status`,`t3`.`address` FROM `test` as t1" +
		" LEFT JOIN `test2` as `t2` ON `t1`.`class` = `t2`.`class`" +
		" INNER JOIN `test3` as t3 ON `t1`.`email` = `t3`.`email` WHERE `t1`.`age` >= ?" +
		" GROUP BY `t1`.`age` HAVING COUNT(`t1`.`age`) > ? ORDER BY `t1`.`age` DESC LIMIT ? OFFSET ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetQueryParams()

	if params[0].(int) != 18 {
		t.Error("params gen err")
	}
}
func TestBuilderInsert(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Insert([]string{"`name`", "`age`"}, "jack", 18).
		GetInsert()
	if err != nil {
		t.Error(err)
	}

	expect := "INSERT INTO `test` (`name`,`age`) VALUES (?,?)"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetInsertParams()

	if params[0].(string) != "jack" ||
		params[1].(int) != 18 {
		t.Error("params gen err")
	}
}

func TestBuilderUpdate(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Update([]string{"`name`", "`age`"}, "jack", 18).
		Where("`id`", "=", 11).
		GetUpdate()
	if err != nil {
		t.Error(err)
	}

	expect := "UPDATE `test` SET `name` = ?,`age` = ? WHERE `id` = ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetUpdateParams()

	if params[0].(string) != "jack" ||
		params[1].(int) != 18 ||
		params[2].(int) != 11 {
		t.Error("params gen err")
	}
}

func TestBuilderDelete(t *testing.T) {
	sb := NewBuilder()

	sql, err := sb.Table("`test`").
		Where("`id`", "=", 11).
		GetDelete()
	if err != nil {
		t.Error(err)
	}

	expect := "DELETE FROM `test` WHERE `id` = ?"
	if sql != expect {
		t.Error("sql gen err")
	}

	params := sb.GetDeleteParams()

	if params[0].(int) != 11 {
		t.Error("params gen err")
	}
}

func TestGenPlaceholders(t *testing.T) {
	pss := []string{
		GenPlaceholders(5),
		GenPlaceholders(3),
		GenPlaceholders(1),
		GenPlaceholders(0),
	}
	results := []string{
		"?,?,?,?,?",
		"?,?,?",
		"?",
		"",
	}

	for k, ps := range pss {
		if ps != results[k] {
			t.Errorf("%s not equal to %s\n", ps, results[k])
		}
	}

}

func BenchmarkQuery(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewBuilder()
		_, err := sb.Table("`test`").
			Select("`name`", "`age`", "`email`").
			Where("`name`", "=", "jack").
			Where("`age`", ">=", 18).
			OrderBy("DESC", "`age`").
			Limit(1, 10).
			GetQuery()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSelect(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewBuilder()
		_, err := sb.Table("`test`").
			Select("`name`", "`age`", "`email`").
			GetQuery()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWhere(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewBuilder()
		_, err := sb.Table("`test`").
			Where("`age`", ">=", 18).
			GetQuery()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWhereIn(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewBuilder()
		_, err := sb.Table("`test`").
			WhereIn("`age`", 18, 19, 20, 31, 22, 33, 24, 45).
			GetQuery()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWhereRaw(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewBuilder()
		_, err := sb.Table("`test`").
			WhereRaw("`age` >= ?", 18).
			GetQuery()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGroupBy(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewBuilder()
		_, err := sb.Table("`test`").
			GroupBy("`age`").
			GetQuery()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkHaving(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewBuilder()
		_, err := sb.Table("test").
			Select("`email`", "`class`", "COUNT(*) as `ct`").
			GroupBy("`email`", "`class`").
			Having("`ct`", ">", "2").
			GetQuery()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkHavingRaw(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewBuilder()
		_, err := sb.Table("`test`").
			Select("`email`", "`class`", "COUNT(*)").
			GroupBy("`email`", "`class`").
			HavingRaw("COUNT(*) > 2").
			GetQuery()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsert(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewBuilder()
		_, err := sb.Table("`test`").
			Insert([]string{"`name`", "`class`"}, "bob", "2-3").
			GetInsert()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUpdate(b *testing.B) {

	for i := 0; i < b.N; i++ {
		sb := NewBuilder()
		_, err := sb.Table("`test`").
			Update([]string{"`name`", "`class`"}, "bob", "2-3").
			GetUpdate()
		if err != nil {
			b.Fatal(err)
		}
	}
}
