## go-nsx

This is an NSX client library for Go. If there is no client API, there will be a type.

```powershell
foreach ($file in gci .\templates\) { $OUTFILE = $file.Name.Replace(".xsd.xml", "");$TYPES =
E:\Programming\Go\bin\Dley -G .\templates\$file; Write-Output "package go_nsx" $OFS $TYPES $OFS | Out-File
.\$OUTFILE.go}; foreach ($file in gci .\) { gofmt -w $file . }
```

### Notes

Most of this code was autogenerated, so it will need tests and stuff like that. Also, APIs may not be under the right
 section, feel free to put in a PR for it.