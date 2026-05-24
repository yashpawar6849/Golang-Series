Golang Series
=============

A small, lesson-based repository for learning Go (Golang). Each lesson is a self-contained folder with a `main.go` you can run to explore language features and examples.

**Repository Structure**
- `Lesson 1/` : Hello From Go Program
- `Lesson 2/` : Data Types in Go
- `Lesson 3/` : Integers (Signed, Unsigned, Bytes, and Runes)

**Prerequisites**
- Go installed (recommended Go 1.20+). Verify with: `go version`

**Run a lesson**
From the repository root, run a lesson by changing into its directory and running the module:

```
cd "Lesson 1"
go run .
```

Or run directly from the workspace root:

```
cd "Golang Series"
go run ./Lesson\ 1
```

Replace `Lesson 1` with `Lesson 2` or `Lesson 3` as needed.

**Build a binary**

```
cd "Lesson 3"
go build -o lesson3
./lesson3
```

**Notes**
- Each lesson is intentionally small and focused. Feel free to modify the `main.go` files while experimenting.
- If you add dependencies, use `go mod init` and `go mod tidy` inside the lesson directory (or at repo root) as appropriate.

**Contributing**
- Contributions and improvements are welcome. Open an issue or submit a PR with additions or clearer explanations.

Enjoy learning Go!
