PACKAGE

package graphite
    import "/home/bhawn/graphite"


TYPES

type Graphite struct {
    // contains filtered or unexported fields
}
    This defined the Graphite structure which contains the necessary bits
    for a functioning object.

func NewGraphite(endpoint string, interval, timeout time.Duration, retries int, debug bool) (*Graphite, error)
    Function to build and return a Graphite object

    Example g = graphite.NewGraphite("Localhost:2003", 10 * time.Second, 2 *
    time.Second, 3, false) returns a Graphite object which writes to the
    Carbon server running on the localhost on port 2003 at 10 second
    intervals with a 2 second timeout to writes and no debugging. It will
    attempt to reconnect to the Carbon server 3 times before failing out
    entirely.

func (g *Graphite) Insert(name string, val interface{})
    Function to insert key:value pairs into the channel

func (g *Graphite) PostAll()
    Wrapper function for our writer that cycles through any available data
    pairs and then sends them to the writer

func (g *Graphite) PostOne(name string, val interface{}) error
    write to our carbon server, if the connection is broken, we will retry
    for the number of retries specified before failing out entirely It is
    acceptable to call PostOne directly if your program will exit on
    completion so you do not need to hang around waiting for the loop to
    execute

func (g *Graphite) Shutdown()
    Function to shutdown the Graphite object cleanly


