# howsmyhttp

This is the source for [howsmyhttp.com](https://howsmyhttp.com).

How's My HTTP tells you which kind of HTTP your client supports. I built this in an effort to practice [Go](https://golang.org). Inspired by the eminently more useful [howsmyssl](https://github.com/jmhodges/howsmyssl).

## Development

HTML templates, CSS, and JavaScript live in `src/`. `sh/start.sh` watches `src/` for changes and recompiles the project when the templates change.

## Scripts
<table>
   <tr>
     <td><strong><code>build.sh</code></strong></td>
     <td>Runs go build and generates the <code>howsmyhttp</code> binary</td>
    </tr>
    <tr>
      <td><strong><code>start.sh</code></strong></td>
      <td>Starts the server, watches <code>src/</code> and recompiles when the templates change</td>
    </tr>
</table>

## Dependencies

- [`justrun`](https://github.com/jmhodges/justrun)
  - [Download a pre-built binary](http://projects.somethingsimilar.com/justrun/downloads)
  - Build from source: `go get github.com/jmhodges/justrun`