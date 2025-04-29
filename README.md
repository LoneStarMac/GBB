# Git Build Booster (GBB) 🚀

**Wraps Homebrew and CMake to fix, patch, and build almost any open-source Git project on macOS with one command!**

## Quickstart
Clone your project normally:
```bash
git clone https://github.com/dosbox-staging/dosbox-staging.git
cd dosbox-staging
```

Then run:

```bash
git clone https://github.com/lonestarmac/gbb.git
./gbb/gbb.sh -DDEBUG=ON -DCMAKE_BUILD_TYPE=Debug
```

GBB will:
  -  Detect missing libraries
  -  Fix CMake errors automatically
  -  Apply optional patches
  -  Build your project
  -  Capture errors if needed

Why use GBB?

Building from source on macOS should just work*. GBB bridges the gap between upstream instructions and Homebrew paths and dependencies.