# Serialization

This section discuss techniques to serialize data.

## Abstract Syntax Notation One (ASN.1)

Abstract Syntax Notation One (ASN.1) is a standard interface description language for defining data structures that can be serialized and deserialized in a cross-platform way. It is broadly used in telecommunications and computer networking, and especially in cryptography.

### Working examples

* [Basic example](../internal/asnser/asn1_test.go)

### References

* [Introduction to ASN.1](https://www.itu.int/en/ITU-T/asn1/Pages/introduction.aspx)
* [OSS Nokalva - ASN](https://www.oss.com/resources/resources.html)
* [ASN1 Simple types](https://www.obj-sys.com/asn1tutorial/node10.html)
* [A Layman's Guide to a Subset of ASN.1, BER, and DER](http://luca.ntop.org/Teaching/Appunti/asn1.html)

##  Base64

In computer programming, Base64 is a group of binary-to-text encoding schemes that represent binary data (more specifically, a sequence of 8-bit bytes) in sequences of 24 bits that can be represented by four 6-bit Base64 digits[wiki](https://en.wikipedia.org/wiki/Base64)).

### Working examples

* [Basic example](../internal/base64ser/base64_test.go)

### References

* [The Base16, Base32, and Base64 Data Encodings](https://datatracker.ietf.org/doc/html/rfc4648)

## Ini File

A text based configuration file comprising of key value pair

### Working examples

* [Basic example](../cmd/inifile/main.go)

## The Concise Binary Object Representation (CBOR)

The Concise Binary Object Representation (CBOR) -- RFC 8949-- is a data format whose design goals include the possibility of extremely small code size, fairly small message size, and extensibility without the need for version negotiation[CBOR](https://cbor.io/).

### Working examples

* [Basic example](../internal/cborser/cbor_test.go)

## Gob

This is a Go-specific data package for communicating between two servers written in Go.

### Working examples

* [Basic example](../internal/gobser/gob_test.go)

### References

* [Gobs of data](https://go.dev/blog/gob)


## Tom's Obvious Minimal Language (TOML)

TOML is a minimal configuration file format that's easy to read due to obvious semantics. TOML is designed to map unambiguously to a hash table.

### Working examples

* [Basic example](../internal/tomlser/toml_test.go)

### References

* [Official Documentation](https://toml.io/en/)

## YAML Ain’t Markup Language (YAML)

YAML (a recursive acronym for “YAML Ain’t Markup Language”) is a data serialization language based on a the use of indentation. It is intended to be human readable.

### Working examples

* [Basic example](../internal/ymlser/yaml_test.go)

### References

* [Official documentation](https://yaml.org/spec/1.2.2/)