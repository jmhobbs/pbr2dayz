[![Go Reference](https://pkg.go.dev/badge/github.com/jmhobbs/pbr2dayz.svg)](https://pkg.go.dev/github.com/jmhobbs/pbr2dayz)
[![Lints](https://github.com/jmhobbs/pbr2dayz/actions/workflows/lint.yml/badge.svg)](https://github.com/jmhobbs/pbr2dayz/actions/workflows/lint.yml)
[![Tests](https://github.com/jmhobbs/pbr2dayz/actions/workflows/test.yml/badge.svg)](https://github.com/jmhobbs/pbr2dayz/actions/workflows/test.yml)
[![codecov](https://codecov.io/github/jmhobbs/pbr2dayz/graph/badge.svg?token=sB2axgNro5)](https://codecov.io/github/jmhobbs/pbr2dayz)

# PBR2DayZ

Convert PBR textures to DayZ format, and DayZ textures back to PBR.

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

A matching reverse CLI is provided in `cmd/dayz2pbr`:

```
usage: dayz2pbr <co> <nohq> <as> <smdi>
```

Example:

```
$ ./dayz2pbr ducky_base_co.png ducky_base_nohq.png ducky_base_as.png ducky_base_smdi.png
$ ls -lart ducky_base*
total 8016
-rw-r--r--@   1 jmhobbs  staff  138239 Apr 12 15:43 ducky_base_basecolor.png
-rw-r--r--@   1 jmhobbs  staff  443456 Apr 12 15:43 ducky_base_normal.png
-rw-r--r--@   1 jmhobbs  staff  135524 Apr 12 15:43 ducky_base_ao.png
-rw-r--r--@   1 jmhobbs  staff  223530 Apr 12 15:43 ducky_base_metallic.png
-rw-r--r--@   1 jmhobbs  staff  223530 Apr 12 15:43 ducky_base_roughness.png
```
