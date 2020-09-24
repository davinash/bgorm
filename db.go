package bgorm

// DB contains information for the current db connection
type DB struct {
}

// Open initialize the database connection, need to import driver first e.g.,
//     import _ "github.com/davinash/bgorm/driver/cassandra"
//     func main() {
//       db, err := gorm.Open("cassandra", "connection string")
//     }
// BGORM has wrapped some drivers, for easier to remember driver's import path, so you could import the cassandra driver with
//    import _ "github.com/davinash/bgorm/driver/cassandra"

func Open(dialect string, args string) (*DB, error) {
	return nil, nil
}

type OperationType string

const (
	Insert OperationType = "Insert"
	Read   OperationType = "Read"
	Delete OperationType = "Delete"
)

type CallBackLevel string

const (
	OnDatabase CallBackLevel = "OnDatabaseLevel"
	OnTable    CallBackLevel = "OnTableLevel"
)

// RegisterCallBack Get called after the operation
func (d *DB) RegisterCallBack(level CallBackLevel, operationType OperationType) {
}
