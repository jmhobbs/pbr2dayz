[![Go Reference](https://pkg.go.dev/badge/github.com/jmhobbs/pbr2dayz.svg)](https://pkg.go.dev/github.com/jmhobbs/pbr2dayz)
[![Lints](https://github.com/jmhobbs/pbr2dayz/actions/workflows/lint.yml/badge.svg)](https://github.com/jmhobbs/pbr2dayz/actions/workflows/lint.yml)
[![Tests](https://github.com/jmhobbs/pbr2dayz/actions/workflows/test.yml/badge.svg)](https://github.com/jmhobbs/pbr2dayz/actions/workflows/test.yml)
[![codecov](https://codecov.io/github/jmhobbs/pbr2dayz/graph/badge.svg?token=sB2axgNro5)](https://codecov.io/github/jmhobbs/pbr2dayz)

# PBR2DayZ

Convert PBR textures to DayZ format.

## Usage

A basic CLI tool is provided in `cmd/pbr2dayz`:

```
usage: pbr2dayz <basecolor> <normal> <ao> <metallic> <roughness>
```

Example:

```
$ ./pbr2dayz ducky_base.png ducky_nor.png ducky_ao.png ducky_metal.png ducky_rough.png
$ ls -lart ducky_base*
total 4008
-rw-r--r--@   1 jmhobbs  staff  138239 Apr 12 15:43 ducky_base.png
-rw-r--r--@   1 jmhobbs  staff  108872 Apr 12 15:43 ducky_base_co.png
-rw-r--r--@   1 jmhobbs  staff  443456 Apr 12 15:43 ducky_base_nohq.png
-rw-r--r--@   1 jmhobbs  staff  135524 Apr 12 15:43 ducky_base_as.png
-rw-r--r--@   1 jmhobbs  staff  223530 Apr 12 15:43 ducky_base_smdi.png
```
