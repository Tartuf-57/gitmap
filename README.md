# gitmap

A command-line tool built in Go that displays a Git contribution
heatmap for a selected author email across local repositories.

## Installation
From the project directory:

```powershell
go install .
```

## Usage

### Add/update repositories to the scan list

Pass a folder containing several repos:

```powershell
gitmap -add "C:\dev"
```
### Display contribution statistics

```powershell
gitmap.exe -email "your_email"
```
