# General Documentation

## Variables

All **variables** and **parameters** have to be `camelCase`. Parameters are preferred to be kept short, unless it is not possible.  
All **fields** in structs or elsewhere are Exported and should be `PascalCase`.  
In the Database all **column names** are written in `camelCase`.  

## Functions

All Functions should be either in `PascalCase` if exported, or in `camelCase` in case they are not to be exported.  
All Methods have pointer recivers (`func (x *MyType) MyFunc()`), as it is faster and the Fields of said Recivers can be changed.
