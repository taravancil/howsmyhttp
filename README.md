# howsmyhttp

This is the source for [howsmyhttp.com](https://howsmyhttp.com).

How's My HTTP tells you which kind of HTTP your client supports. I built this in an effort to practice [Go](https://golang.org). Inspired by the eminently more useful [howsmyssl](https://github.com/jmhodges/howsmyssl).

## Development

HTML templates, CSS, and JavaScript live in `src/`. `sh/start.sh` watches `src/` for changes and recompiles the project when the templates change.

## Scripts
All scripts live in `sh/`.

<table>
   <tr>
     <td><strong><code>build.sh</code></strong></td>
     <td>Runs go build and generates the <code>howsmyhttp</code> binary</td>
    </tr>
    <tr>
      <td><strong><code>start.sh</code></strong></td>
      <td>Starts the server, watches <code>src/</code> and recompiles when the templates change</td>
    </tr>
    <tr>
      <td><strong><code>subset-fonts.sh</code></strong></td >
      <td>Generates subset font files from <code>src/fonts/</code> and outputs them to <code>static/fonts/</code></td>
    </tr>
</table>

## Dependencies

- [`justrun`](https://github.com/jmhodges/justrun)
  - [Download a pre-built binary](http://projects.somethingsimilar.com/justrun/downloads)
  - Build from source
    ```bash
    go get github.com/jmhodges/justrun
    ```

## Typography

`howsmyhttp` uses Verdana, a web-safe font, for all body copy. [Nunito](https://fonts.google.com/specimen/Nunito) is used for headings and decorative text. To reduce the size of font files, all loaded font files are subset to only include the glyphs in use on [howsmyhttp.com](https://../howsmyhttp.com)

### Regenerating Subset Font Files

If you add text with glyphs not yet included in the subset font files, you'll need to regenerate them with [`glyphhanger`](https://github.com/filamentgroup/glyphhanger).

Install glyphanger with `npm`:
```bash
npm i -g glyphhanger
```

*`glyphhanger` also requires `pyftsubset` and `brotli` for generating WOFF2 files. See the [installation instructions](https://github.com/filamentgroup/glyphhanger#prerequisite-pyftsubset) if you don't already have these installed.*

With the local server running, run `sh/subset-fonts.sh`. The generated font files will be moved to `static/fonts`. Copy the output `@font-face` declarations into `src/css/fonts.css`.