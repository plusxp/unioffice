//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package tempstorage ;import _a "io";

// Open returns tempstorage File object by name.
func Open (path string )(File ,error ){return _b .Open (path )};type storage interface{Open (_gg string )(File ,error );TempFile (_ga ,_d string )(File ,error );TempDir (_da string )(string ,error );RemoveAll (_gc string )error ;Add (_f string )error ;};

// Add reads a file from a disk and adds it to the storage.
func Add (path string )error {return _b .Add (path )};

// File is a representation of a storage file
// with Read, Write, Close and Name methods identical to os.File.
type File interface{_a .Reader ;_a .Writer ;_a .Closer ;Name ()string ;};

// TempFile creates new empty file in the storage and returns it.
func TempFile (dir ,pattern string )(File ,error ){return _b .TempFile (dir ,pattern )};

// TempDir creates a name for a new temp directory using a pattern argument.
func TempDir (pattern string )(string ,error ){return _b .TempDir (pattern )};

// SetAsStorage changes temporary storage to newStorage.
func SetAsStorage (newStorage storage ){_b =newStorage };var _b storage ;

// RemoveAll removes all files according to the dir argument prefix.
func RemoveAll (dir string )error {return _b .RemoveAll (dir )};