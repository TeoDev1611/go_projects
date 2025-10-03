# 🔐 GoPass - Secure Password Generator

A cryptographically secure random password generator developed in Go, featuring an interactive terminal interface.

## ✨ Features

- **Cryptographically secure generation**: Uses `crypto/rand` to ensure high-quality randomness
- **Interactive interface**: Intuitive terminal-based menu powered by the Survey library
- **Customizable**: Configure exactly which character types to include
- **Ambiguous character exclusion**: Option to avoid visually confusing characters (like `1`, `l`, `I`, `0`, `O`)
- **No statistical bias**: Implements an algorithm that eliminates modular bias in random generation

## 🚀 Installation

### Prerequisites

- Go 1.16 or higher installed on your system

### Steps

1. Clone the repository:
```bash
git clone <repository-url>
cd gopass
```

2. Install dependencies:
```bash
go get github.com/AlecAivazis/survey/v2
```

3. Build the project:
```bash
go build -o gopass
```

## 💻 Usage

Run the program:

```bash
./gopass
```

### Configuration Options

The program will present an interactive menu where you can select:

- ✅ **Lowercase (a-z)** - Lowercase letters
- ✅ **Uppercase (A-Z)** - Uppercase letters
- ✅ **Digits (0-9)** - Numbers from 0 to 9
- ⬜ **Symbols (!@#$%...)** - Special characters
- ⬜ **Exclude ambiguous characters** - Avoids characters like `1lI0Oo8B` that can be confused

*By default, lowercase, uppercase, and digits are selected.*

### Usage Example

```
 ::::::::  :::::::: :::::::::    :::     ::::::::  ::::::::  
:+:    :+::+:    :+::+:    :+: :+: :+:  :+:    :+::+:    :+: 
+:+       +:+    +:++:+    +:++:+   +:+ +:+       +:+        
:#:       +#+    +:++#++:++#++#++:++#++:+#++:++#+++#++:++#++ 
+#+   +#+#+#+    +#++#+      +#+     +#+       +#+       +#+ 
#+#    #+##+#    #+##+#      #+#     #+##+#    #+##+#    #+# 
 ########  ######## ###      ###     ### ########  ########  

? Select options for the character pool:
  ◉ Use lowercase (a-z)
  ◉ Use uppercase (A-Z)
  ◉ Use digits (0-9)
  ◯ Use symbols (!@#$%...)
  ◉ Exclude ambiguous characters (l,1,O,0...)

📊 Selection summary:
  Lowercase: true
  Uppercase: true
  Digits: true
  Symbols: false
  Exclude ambiguous: true

🎯 Generated pool (54 characters):
abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789

? Enter your password length: 20

Generated password: mK7p3nRx9TqW5zFh4YvG
```

## 🏗️ Code Structure

### Main Functions

- **`NuevosConjuntos()`**: Defines available character sets
- **`CrearPool()`**: Generates the character pool based on selected options
- **`CrearPwsd()`**: Generates the random password
- **`indiceAleatorioSeguro()`**: Implements unbiased index generation without modular bias
- **`menuPool()`**: Interactive interface for option selection
- **`pedirNumero()`**: Requests the desired password length

## 🔒 Security

This generator implements the following security measures:

1. **Use of `crypto/rand`**: Cryptographically secure randomness source provided by the operating system
2. **Modular bias elimination**: The algorithm rejects values that could introduce statistical bias in the distribution
3. **No storage**: Passwords are not saved to any file or history

## 📋 Dependencies

- [github.com/AlecAivazis/survey/v2](https://github.com/AlecAivazis/survey) - For the interactive terminal interface

## 🤝 Contributing

Contributions are welcome. Please:

1. Fork the project
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is under the MIT License. See the `LICENSE` file for more details.

## 👨‍💻 Author

Developed with ❤️ using Go in Ecuador