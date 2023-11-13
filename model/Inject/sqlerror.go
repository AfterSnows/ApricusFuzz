package Inject

import (
	"log"
	"regexp"
)

var sqlErrors = map[string]string{
	"SQL syntax":                                              "MYSQL",
	"syntax to use near":                                      "MYSQL",
	"MySQLSyntaxErrorException":                               "MYSQL",
	"valid MySQL result":                                      "MYSQL",
	"SQL syntax.*?MySQL":                                      "MYSQL",
	"Warning.*?mysql_":                                        "MYSQL",
	"MySqlException \\(0x":                                    "MYSQL",
	"PostgreSQL.*?ERROR":                                      "PostgreSQL",
	"Warning.*?\\Wpg_":                                        "PostgreSQL",
	"valid PostgreSQL result":                                 "PostgreSQL",
	"Npgsql\\.":                                               "PostgreSQL",
	"PG::SyntaxError:":                                        "PostgreSQL",
	"org\\.postgresql\\.util\\.PSQLException":                 "PostgreSQL",
	"ERROR:\\s\\ssyntax error at or near":                     "PostgreSQL",
	"Driver.*? SQL[-\\_ ]*Server":                             "Microsoft SQL Server",
	"OLE DB.*? SQL Server":                                    "Microsoft SQL Server",
	"SQL Server[^<&quot;]+Driver":                             "Microsoft SQL Server",
	"Warning.*?(mssql|sqlsrv)_":                               "Microsoft SQL Server",
	"SQL Server[^<&quot;]+[0-9a-fA-F]{8}":                     "Microsoft SQL Server",
	"System\\.Data\\.SqlClient\\.SqlException":                "Microsoft SQL Server",
	"(?s)Exception.*?\\WRoadhouse\\.Cms\\.":                   "Microsoft SQL Server",
	"Microsoft SQL Native Client error '[0-9a-fA-F]{8}":       "Microsoft SQL Server",
	"com\\.microsoft\\.sqlserver\\.jdbc\\.SQLServerException": "Microsoft SQL Server",
	"ODBC SQL Server Driver":                                  "Microsoft SQL Server",
	"ODBC Driver \\d+ for SQL Server":                         "Microsoft SQL Server",
	"macromedia\\.jdbc\\.sqlserver":                           "Microsoft SQL Server",
	"com\\.jnetdirect\\.jsql":                                 "Microsoft SQL Server",
	"SQLSrvException":                                         "Microsoft SQL Server",
	"Microsoft Access (\\d+ )?Driver":                         "Microsoft Access",
	"ODBC Microsoft Access":                                   "Microsoft Access",
	"Syntax error \\(missing operator\\) in query expression": "Microsoft Access",
	"ORA-\\d{5}":                                              "Oracle",
	"Oracle error":                                            "Oracle",
	"Oracle.*?Driver":                                         "Oracle",
	"Warning.*?\\Woci_":                                       "Oracle",
	"Warning.*?\\Wora_":                                       "Oracle",
	"oracle\\.jdbc\\.driver":                                  "Oracle",
	"quoted string not properly terminated":                   "Oracle",
	"SQL command not properly ended":                          "Oracle",
	"DB2 SQL error":                                           "CLI Driver.*?DB2",
	"db2_\\w+\\(":                                             "CLI Driver.*?DB2",
	"SQLSTATE.+SQLCODE":                                       "CLI Driver.*?DB2",
	"check the manual that corresponds to your (MySQL|MariaDB) server version": "MYSQL",
	"Unknown column '[^ ]+' in 'field list'":                                   "MYSQL",
	"MySqlClient\\.":                                                           "MYSQL",
	"com\\.mysql\\.jdbc\\.exceptions":                                          "MYSQL",
	"Zend_Db_Statement_Mysqli_Exception":                                       "MYSQL",
	"Access Database Engine":                                                   "Microsoft Access",
	"JET Database Engine":                                                      "Microsoft Access",
	"Microsoft Access Driver":                                                  "Microsoft Access",
	"SQLServerException":                                                       "Microsoft SQL Server",
	"SqlException":                                                             "Microsoft SQL Server",
	"SQLServer JDBC Driver":                                                    "Microsoft SQL Server",
	"Incorrect syntax":                                                         "Microsoft SQL Server",
	"MySQL Query fail":                                                         "MYSQL",
	"Unknown column.*?order clause":                                            "MYSQL",
}

func ExtractErrorFromBody(responseBody string) bool {
	for pattern, _ := range sqlErrors {
		rx := regexp.MustCompile(pattern)
		if rx.MatchString(responseBody) {
			log.Print("匹配错误信息成功:" + pattern)
			return true
		}
	}
	return false
}
