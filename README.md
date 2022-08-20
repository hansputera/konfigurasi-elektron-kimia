## Konfigurasi Elektron

### A. Kulit
Terdiri dari kulit `K L M N, dst`. Maksimum elektron pada tiap kulit adalah `2n^2`

**Contoh:**
```md
* Atom          K   L   M   N   O   P   Q

* 3^Li          2   1
* 9^F           2   7
* 11^Na         2   8   1
* 20^Ca         2   8   8   2
* 55^Cs         2   8   18  18  8   1
* 53^I          2   8   18  18  7
* 86^Rn         2   8   18  32  18  8
* 87^Fr         2   8   18  32  18  8   1
```

**Contoh implementasi kode:**
```go
atom := parsers.GetAtom("3^Li")
// atom.Unsur = Li
// atom.Massa = 0 (karena tidak didefinisikan)
// atom.Nomor = 3

kulits := functions.KulitAtom(atom)
// kulits = {K L M N O P Q}
// kulits = {2 1 0 0 0 0 0}
```