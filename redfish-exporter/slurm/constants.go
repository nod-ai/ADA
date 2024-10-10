package slurm

// Method names used in the Slurm API
const (
	methodPostNode = "PostNode"
	methodPing     = "Ping"
)

// withExecuteSuffix appends "Execute" to the given method name.
// For example, "PostNode" becomes "PostNodeExecute".
func withExecuteSuffix(method string) string {
	return method + "Execute"
}
