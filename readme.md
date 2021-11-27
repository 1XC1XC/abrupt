## abrupt
	go get -d github.com/1XC1XC/abrupt
	
Import
```golang
// WIP
```

Random
```golang
// Numbers
Rand.Int(1,100) // Random Number 1 - 100
Rand.Int(10) // Random Number 0 - 10
Rand.Int() // Random Number 0 - 5
Rand.Float(1,100)  // Random Decimal Number 1 - 100
Rand.Float(10)  // Random Decimal Number 0 - 10
Rand.Float()  // Random Decimal Number 0 - 5

// Strings
Rand.Str(10) // Random 10 Characters [All Characters]
Rand.Str(10, true) // Random 10 Characters [Uppercase/Lowercase Letters]
Rand.Str() // Random 5 Characters [All Characters]
Rand.Str(true) // Random 5 Characters [Uppercase/Lowercase Letters]

// Boolean
Rand.Bool() // Random Booleans (true, false)

// Arrays
Rand.Array([]int{1, 2, 3}) // Random Number Element (1, 2, 3)  
Rand.Array([]string{"a", "b", "c"}) // Random String Element (a, b, c)
Rand.Array([]interface{}{1, "b", true}) // Random Element (1, b, true)

/* Rand.Array: Creating Variable Practices 
	Number := Rand.Array([]int{1, 2, 3})
    or
	var Letter interface{} = Rand.Array([]string{"a", "b", "c"})
*/
```

Crypto
```golang
Crypto.Base64.Encode("Hello World!") // Base64 Encoding (YXNkZg==)
Crypto.Base32.Encode("Hello World!") // Base32 Encoding (JBSWY3DPEBLW64TMMQQQ====)
Crypto.Hex.Encode("Hello World!") // Base16 Encoding (48656c6c6f20576f726c6421)

Crypto.SHA512("Hello World!") // SHA512 Hash (861844d6704e8573fec34d...)
Crypto.SHA256("Hello World!") // SHA256 Hash (7f83b1657ff1fc53b92dc1...)
Crypto.MD5("Hello World!") // MD5 Hash (ed076287532e86365e841e92bfc5...)
```

File
```golang
// Create, Read functions return false if (folder/file) are absent
File.Create("./Directory") // New Folder
File.Create("./Directory/A.txt", "Hello") // New File 
File.Create("./Directory/B.txt", "World") // New File


File.Read("./Directory") // Directory Array ([A.txt B.txt])
File.Read("./Directory/A.txt") // File String (Hello)
File.Read("./Directory/B.txt") // File String (World)

File.Exists("./Directory") // (true) true/false depending on the presence in directory
File.Exists("./Directory", true) // (folder) while second parameter is true returns the type
File.Exists("./Directory/A.txt", true) // (file)

File.Remove("./Directory/A.txt") // (true) output is true after running accurately
File.Remove("./Directory") // (true)
File.Remove("./File") // (false) (assuming there is no file/directory present there)
```

String
```golang 
String.Reverse("Hello World!") // (!dlroW olleH)

String.Comma("10000") // (10,000)
String.Comma(10000) // (10,000)
String.Comma("10000", "$") // ($10,000)
String.Comma(10000, "$") // ($10,000)
```