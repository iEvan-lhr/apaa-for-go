package structs

import (
	_ "github.com/go-sql-driver/mysql"
	tools "github.com/iEvan-lhr/exciting-tool"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"sync"
	"xorm.io/xorm"
)

type Query struct {
	Sql    string
	Model  any
	Result chan []map[string]interface{}
}

type Where struct {
	Key   string
	Value string
}

type QueryStruct struct {
	Type   bool
	Where  []Where
	Count  bool
	Exists bool
	Struct any
}

type DBConn struct {
	DB         *xorm.Engine
	execChan   chan string
	execMap    sync.Map
	lock       bool
	queryChan  chan *Query
	queryMap   sync.Map
	resultsMap sync.Map
}

func (d *DBConn) GetConn(mission chan *anything.Mission, data []any) {
	if d.DB == nil {
		d.DB = (<-anything.DoChanN("InitDB", data)).Pursuit[0].(*xorm.Engine)
	}
	mission <- &anything.Mission{Pursuit: []any{d.DB}}
}

func (d *DBConn) InitDB(mission chan *anything.Mission, data []any) {
	engine := tools.ReturnValueByTwo(xorm.NewEngine(data[0].(string), data[1].(string))).(*xorm.Engine)
	d.DB = engine
	go d.initQuerySql()
	go d.initExecSql()
	mission <- &anything.Mission{Pursuit: []any{engine}}
}

func (d *DBConn) initExecSql() {
	d.execChan = make(chan string, 20)
	d.execMap = sync.Map{}
	for i := 0; i < 20; i++ {
		d.execMap.Store(i, make(chan string, 2))
		go func(index int) {
			te, _ := d.execMap.Load(index)
			ch := te.(chan string)
			for {
				exec := <-ch
				if exec == "EXIT" {
					return
				} else {
					_, err := d.DB.Exec(exec)
					anything.ErrorDontExit(err)
				}
			}
		}(i)
	}
	k := 0
	for {
		insert := <-d.execChan
		if insert == "EXIT" {
			for i := 0; i < 20; i++ {
				ch, _ := d.execMap.Load(i)
				ch.(chan string) <- "EXIT"
			}
			return
		} else {
			if k < 20 {
				ch, _ := d.execMap.Load(k)
				ch.(chan string) <- insert
				k++
			} else {
				k = 0
				ch, _ := d.execMap.Load(k)
				ch.(chan string) <- insert
			}
		}
	}
}

func (d *DBConn) initQuerySql() {
	d.queryMap = sync.Map{}
	d.queryChan = make(chan *Query, 15)
	for i := 0; i < 10; i++ {
		d.queryMap.Store(i, make(chan *Query, 2))
		go func(index int) {
			te, _ := d.queryMap.Load(index)
			ch := te.(chan *Query)
			for {
				querySql := <-ch
				if result, err := d.DB.QueryInterface(querySql.Sql); err == nil {
					querySql.Result <- result
				} else {
					anything.ErrorDontExit(err)
				}
			}
		}(i)
	}
	k := 0
	for {
		queryS := <-d.queryChan
		if k < 10 {
			ch, _ := d.queryMap.Load(k)
			ch.(chan *Query) <- queryS
			k++
		} else {
			k = 0
			ch, _ := d.queryMap.Load(k)
			ch.(chan *Query) <- queryS
		}

	}
}

func (d *DBConn) QuerySql(data []any) {
	d.queryChan <- data[0].(*Query)
}

func (d *DBConn) QueryStruct(mission chan *anything.Mission, data []any) {
	session := d.DB.NewSession()
	tools.ExecError(session.Begin())
	queryStruct := data[0].(*QueryStruct)
	if queryStruct.Count {
		count := tools.ReturnValueByTwo(session.Count(queryStruct.Struct)).(int64)
		tools.ExecError(session.Commit())
		mission <- &anything.Mission{Pursuit: []any{count}}
		return
	}
	if queryStruct.Exists {
		exits := tools.ReturnValueByTwo(session.Exist(queryStruct.Struct)).(bool)
		tools.ExecError(session.Commit())
		mission <- &anything.Mission{Pursuit: []any{exits}}
		return
	}
	for i := range queryStruct.Where {
		session.Where(queryStruct.Where[i].Key, queryStruct.Where[i].Value)
	}
	if queryStruct.Type {
		has := tools.ReturnValueByTwo(session.Get(queryStruct.Struct)).(bool)
		tools.ExecError(session.Commit())
		mission <- &anything.Mission{Pursuit: []any{has}}
	} else {
		tools.ExecError(session.Find(queryStruct.Struct))
		tools.ExecError(session.Commit())
		mission <- &anything.Mission{}
	}
}

func (d *DBConn) ExecSql(exec []*tools.String) {
	for i := range exec {
		d.execChan <- exec[i].String()
	}
}

func Find(data, result any) {
	q := &Query{}
	switch data.(type) {
	case string:
		q.Sql = data.(string)
	default:
		q.Sql = tools.Query(data)
	}
	q.Result = make(chan []map[string]interface{})
	anything.OnceSchedule("QuerySql", []any{q})
	res := <-q.Result
	tools.Unmarshal(res, result)
}

func Sure(data any) bool {
	q := &Query{}
	switch data.(type) {
	case string:
		q.Sql = data.(string)
	default:
		q.Sql = tools.Query(data)
	}
	q.Result = make(chan []map[string]interface{})
	anything.OnceSchedule("QuerySql", []any{q})
	res := <-q.Result
	if len(res) > 0 {
		return true
	}
	return false
}

func Check(data any) bool {
	q := &Query{}
	switch data.(type) {
	case string:
		q.Sql = data.(string)
	default:
		q.Sql = tools.Check(data)
	}
	q.Result = make(chan []map[string]interface{})
	anything.OnceSchedule("QuerySql", []any{q})
	res := <-q.Result
	if len(res) > 0 {
		return true
	}
	return false
}

func Save(model any) {
	anything.DoSchedule("ExecSql", []any{tools.Save(model)})
}
