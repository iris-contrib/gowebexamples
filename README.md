## GopherBOOk

GopherBOOk provides easy to understand code snippets on how to get started with web development with the Go programming language using the  [Iris](https://github.com/kataras/iris) web framework.


### Note

This site doesn't contains Iris' 'best ways' neither explains all its features. It's just a simple, practical cookbook for young Go developers!

Developers should look the [Iris web framework](https://github.com/kataras/iris) subfolders for examples and read the official [documentation](https://godoc.org/gopkg.in/kataras/iris.v6).

The site is built using the [Hugo](https://github.com/spf13/hugo) static site generator.
The source codes to the examples can be found in the `content` directory.
When you build the site using the `hugo` command, the output files will be generated into the `public` directory.

### Building

To build this site Hugo and Python (with `pygments` installed) will be required.

```sh
$ git clone https://github.com/iris-contrib/gowebexamples.git
$ cd gowebexamples
$ hugo server
```

### Thanks

Thanks to [Go Web By Example](https://gowebexamples.github.io) for the inspiration to this cookbook.
