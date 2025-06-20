# Go Learning Roadmap - Exercism Exercises

A comprehensive journey through Go programming concepts using Exercism's structured learning path, focusing on simplicity, concurrency, and rapid development.

## 📊 Progress Overview

- **Total Exercises**: 141
- **Completed**: 13 ✅
- **In Progress**: 0 🔄
- **Available**: 113 📋
- **Locked**: 17 🔒
- **Progress**: 9.2%

## 🎯 Current Status

### ✅ Completed Exercises (13)

| Exercise | Concept | Difficulty | Description |
|----------|---------|------------|-------------|
| [Hello World](./hello-world/) | Basics, Packages | Easy | The classical introductory exercise |
| [Gopher's Gorgeous Lasagna](./lasagna/) | Functions, Variables | Easy | Learn about packages, functions, and variables |
| [Annalyn's Infiltration](./annalyns-infiltration/) | Booleans | Easy | Learn about booleans by helping Annalyn |
| [Party Robot](./party-robot/) | Strings | Easy | Learn about strings by programming a party robot |
| [Weather Forecast](./weather-forecast/) | Comments | Easy | Learn about comments for weather forecasting |
| [Blackjack](./blackjack/) | Conditionals | Easy | Learn about conditionals by playing Blackjack |
| [Cars Assemble](./cars-assemble/) | Numbers, Type Conversion | Easy | Learn about numbers analyzing an assembly line |
| [Vehicle Purchase](./vehicle-purchase/) | Comparison, Conditional-ifs | Easy | Learn about comparison while buying vehicles |
| [Card Tricks](./card-tricks/) | Slices | Medium | Learn about slices by doing card tricks |
| [Need For Speed](./need-for-speed/) | Structs | Medium | Learn about structs by racing RC cars |
| [Bird Watcher](./bird-watcher/) | For Loops | Easy | Count birds in your garden with for loops |
| [Lasagna Master](./lasagna-master/) | Variadic Functions | Easy | Dive deeper into Go functions while preparing lasagna |

## 🗺️ Go Learning Path & Concepts

### 🟢 Mastered Concepts
- ✅ **Basics** - Package structure, basic syntax, `main` function
- ✅ **Comments** - Code documentation and inline comments
- ✅ **Numbers** - Integer types, arithmetic operations
- ✅ **Arithmetic Operators** - Mathematical operations (`+`, `-`, `*`, `/`, `%`)
- ✅ **Booleans** - Logical operations and boolean values
- ✅ **Strings** - String manipulation and operations
- ✅ **Conditionals If** - Control flow with if statements
- ✅ **Comparison** - Comparison operators (`==`, `!=`, `<`, `>`, etc.)
- ✅ **String Formatting** - String interpolation with `fmt` package
- ✅ **Packages** - Module organization and imports
- ✅ **Functions** - Function definition, parameters, and return values

### 🟢 Recently Mastered
- ✅ **Slices** - Dynamic arrays, append, and slice operations
- ✅ **Structs** - Custom data types, fields, and methods
- ✅ **For Loops** - Iteration and repetition patterns
- ✅ **Variadic Functions** - Functions with variable arguments (`...`)

### 🟡 Currently Learning
- 🎯 **Ready for next challenge** - All current exercises completed!

### 🔴 Upcoming Core Concepts

#### **Foundation Level (Next 2-3 weeks)**

- 📋 **Maps** - Key-value data structures
- 📋 **Pointers** - Memory addresses and references (`*`, `&`)
- 📋 **Methods** - Functions associated with types

#### **Intermediate Level (Next 1-2 months)**
- 📋 **Interfaces** - Type contracts and polymorphism
- 📋 **Errors** - Error handling patterns and custom errors
- 📋 **Type Definitions** - Custom type creation
- 📋 **Type Assertion** - Runtime type checking
- 📋 **Conditionals Switch** - Switch statement control flow

#### **Advanced Level (Next 2-3 months)**
- 📋 **Goroutines** - Lightweight threads and concurrency
- 📋 **Channels** - Communication between goroutines
- 📋 **Range Iteration** - Iterating over collections
- 📋 **First Class Functions** - Functions as values and closures
- 📋 **Regular Expressions** - Pattern matching with `regexp`
- 📋 **Time** - Date and time manipulation
- 📋 **Floating-point Numbers** - Decimal number operations
- 📋 **Runes** - Unicode character handling
- 📋 **Zero Values** - Default values for types
- 📋 **Stringers** - String representation of custom types
- 📋 **Type Conversion** - Converting between types
- 📋 **Randomness** - Random number generation

## 🎯 Next Priority Exercises

### Easy Exercises - Foundation Building
1. **Two Fer** - Create sentences: "One for X, one for me"
2. **Raindrops** - Convert numbers to strings based on factors
3. **Collatz Conjecture** - Calculate steps to reach 1
4. **Gigasecond** - Calculate time after a gigasecond
5. **Hamming** - Calculate Hamming difference between DNA strands
6. **Scrabble Score** - Compute Scrabble score for words
7. **Leap** - Determine if a year is a leap year
8. **Isogram** - Check if word/phrase is an isogram

### Medium Exercises - Go-Specific Patterns
1. **Strain** - Implement keep/discard operations on collections
2. **Wordy** - Parse and evaluate math word problems
3. **Kindergarten Garden** - Work with plant responsibility diagrams
4. **Tournament** - Process tournament results
5. **Custom Set** - Create custom set data structure

### Hard Exercises - Advanced Go Features
1. **Bank Account** - Concurrent transaction handling
2. **Parallel Letter Frequency** - Goroutines and concurrency
3. **Tree Building** - Refactor tree algorithms

## 📚 Exercise Categories

### 🟢 **Beginner Track (Easy - 42 exercises)**
Foundation concepts, basic Go syntax, and problem-solving fundamentals.

**Key Concepts**: Variables, functions, basic data types, control flow, simple data structures.

### 🟡 **Intermediate Track (Medium - 88 exercises)**
Go-specific features, idiomatic patterns, and moderate complexity algorithms.

**Key Concepts**: Structs, interfaces, methods, error handling, concurrency basics.

### 🔴 **Advanced Track (Hard - 11 exercises)**
Complex problem-solving, performance optimization, and advanced Go features.

**Key Concepts**: Advanced concurrency, complex algorithms, performance tuning.

## 🚀 Getting Started

### Prerequisites
```bash
# Install Go toolchain
go version    # Should be 1.19+

# Install Exercism CLI
# Download from: https://exercism.org/cli-walkthrough

# Verify setup
exercism version
```

### Go Project Structure
```
exercise-name/
├── README.md          # Exercise description
├── go.mod            # Go module file
├── *_test.go         # Test files
├── *.go              # Your solution files
├── HELP.md           # Help and hints
└── .exercism/        # Exercism metadata
```

### Running Exercises
```bash
# Navigate to exercise directory
cd exercism/go/[exercise-name]

# Run tests
go test

# Run tests with verbose output
go test -v

# Run specific test
go test -run TestSpecificFunction

# Check code formatting
go fmt

# Run linter
go vet

# Submit solution
exercism submit *.go
```

### Useful Go Commands
```bash
# Format code
go fmt ./...

# Run all tests in project
go test ./...

# Build executable
go build

# Install dependencies
go mod tidy

# View documentation
go doc fmt.Println

# Check for race conditions
go test -race
```

## 📖 Learning Resources

### Official Go Resources
- [Go Tour](https://go.dev/tour/) - Interactive Go tutorial
- [Go Documentation](https://go.dev/doc/) - Official Go documentation
- [Effective Go](https://go.dev/doc/effective_go) - Best practices guide
- [Go by Example](https://gobyexample.com/) - Practical examples
- [Go Playground](https://go.dev/play/) - Online Go editor

### Community Resources
- [Exercism Go Track](https://exercism.org/tracks/go) - Structured learning path
- [Go Wiki](https://go.dev/wiki/) - Community-maintained wiki
- [Awesome Go](https://awesome-go.com/) - Curated list of Go packages
- [Go Forum](https://forum.golangbridge.org/) - Community discussions
- [r/golang](https://www.reddit.com/r/golang/) - Reddit community

### Books & Advanced Learning
- [The Go Programming Language](https://www.gopl.io/) - Comprehensive guide
- [Go in Action](https://www.manning.com/books/go-in-action) - Practical approach
- [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) - Advanced concurrency

## 🎯 Learning Goals

### Short Term (Next 2 weeks)
- [ ] Complete Need For Speed (Structs)
- [ ] Complete Card Tricks (Slices)
- [ ] Start 3 more Easy exercises
- [ ] Master basic data structures

### Medium Term (Next Month)
- [ ] Complete 10 more exercises
- [ ] Master Methods and Interfaces
- [ ] Start working on Medium difficulty exercises
- [ ] Implement error handling patterns

### Long Term (Next 3 months)
- [ ] Complete 30+ exercises
- [ ] Master goroutines and channels
- [ ] Tackle Hard difficulty exercises
- [ ] Build a substantial Go project

## 🔗 Quick Navigation

- **[← Back to Global Overview](../README.md)** - Multi-language learning dashboard
- **[Rust Learning Track →](../rust/README.md)** - Parallel Rust learning journey

---

## 📝 Go Learning Philosophy

> **"Go is designed to be simple, readable, and efficient. The language's philosophy emphasizes clarity over cleverness, making it an excellent choice for building reliable, maintainable software at scale."**

### **Why Go?**
- **Simplicity** - Clean, readable syntax with minimal keywords
- **Concurrency** - Built-in support with goroutines and channels
- **Performance** - Compiled language with fast execution
- **Productivity** - Fast compilation and excellent tooling
- **Scalability** - Designed for large-scale distributed systems

### **Go's Design Principles**
- **Less is more** - Minimal feature set for maximum clarity
- **Composition over inheritance** - Interfaces and embedding
- **Explicit over implicit** - Clear, obvious code
- **Simple concurrency model** - Goroutines and channels
- **Fast compilation** - Quick feedback loops

### **Learning Approach**
- **Start Simple** - Master the basics before advancing
- **Embrace Conventions** - Follow Go's idiomatic patterns
- **Practice Concurrency** - Learn goroutines and channels early
- **Read Standard Library** - Go's stdlib is excellent learning material
- **Write Tests** - Go has fantastic built-in testing support

### **Common Go Patterns**
- **Error Handling** - Explicit error returns vs exceptions
- **Interface Satisfaction** - Implicit interface implementation
- **Composition** - Embedding types for code reuse
- **Channels** - Communication and synchronization
- **Defer** - Cleanup and resource management

## 📈 Progress Tracking

### **Concept Mastery Checklist**
- [x] **Package System** - Understanding Go modules and imports
- [x] **Basic Syntax** - Variables, functions, control structures
- [x] **Data Types** - Numbers, strings, booleans
- [ ] **Composite Types** - Arrays, slices, maps, structs
- [ ] **Methods & Interfaces** - Go's approach to OOP
- [ ] **Error Handling** - Idiomatic error management
- [ ] **Concurrency** - Goroutines and channels
- [ ] **Testing** - Writing effective tests

### **Skill Development Areas**
- **API Development** - Building web services and APIs
- **CLI Tools** - Command-line application development
- **Concurrent Programming** - Goroutines and channel patterns
- **System Programming** - Lower-level system interaction
- **Microservices** - Distributed system development

**Last Updated**: December 2024  
**Next Review**: Weekly progress updates  
**Focus**: Building strong foundation in Go fundamentals and concurrency