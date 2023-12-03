# SQL Opforu Seeder

this is an alternative to my other mongodb seeder example, showing how you can seed tons of unique data with intact relational references in seconds.

since sql is more strict about what is allowed, the references are strictly defined to allow insertion of only valid references but allows the application layer to take control of the generation meaning network calls for necessary key values are eliminated, reducing the networking penalty that many such seeding tools suffer from.

this is just an example I made mostly for myself and learning purposes. Feel free to use it, take it, modify it, suit it to your needs as you see fit.

I started this back when I didn't really know how to structure GO projects and I still somewhat don't, however most of what you'll probably be interested in is in the `pkg/types/internal.go` file. Take a look at the migrations as well to see how the referneces are validated on the database layer as well as through application logic.

## Usage

Modify the configuration's defaults to connect to the database you want. `pkg/types/store.go`. Ideally the database should be newly created.

Once you've set the configuration you can build and run the binary.

```bash
make run
```

the above will build and run the binary all together. I've included a few helper tasks that might make things simpler.
Take a look inside the _Makefile_ to check them out and use as you wish. 