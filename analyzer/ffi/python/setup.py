"""
Setup script for ANTLR Expression Analyzer Python bindings.
"""
import os
import platform
import subprocess
from pathlib import Path

from setuptools import Extension, find_packages, setup
from setuptools.command.build_ext import build_ext


class GoExtension(Extension):
    """Extension for building Go shared libraries."""
    
    def __init__(self, name, sourcedir=""):
        Extension.__init__(self, name, sources=[])
        self.sourcedir = os.path.abspath(sourcedir)


class GoBuildExt(build_ext):
    """Custom build extension for Go code."""
    
    def run(self):
        for ext in self.extensions:
            if isinstance(ext, GoExtension):
                self.build_go_extension(ext)
        super().run()
    
    def build_go_extension(self, ext):
        """Build Go shared library."""
        print(f"source dir: {ext.sourcedir}")
        extdir = Path(os.path.abspath(
            os.path.dirname(self.get_ext_fullpath(ext.name))
        ))
        
        # Ensure build directory exists
        if not extdir.exists():
            os.makedirs(extdir)
        
        print(f"path: {Path(__file__).absolute()}")

        # Get the Go source directory
        go_src_dir = Path(__file__).parent.parent.parent
        print(f"go_src_dir: {go_src_dir.absolute()}")
        
        # Determine output file name based on platform
        system = platform.system()
        if system == "Darwin":
            output_name = "libanalyzer.dylib"
        elif system == "Linux":
            output_name = "libanalyzer.so"
        elif system == "Windows":
            output_name = "libanalyzer.dll"
        else:
            raise RuntimeError(f"Unsupported platform: {system}")
        
        output_path = os.path.join(extdir, output_name)
        
        # Build command
        cmd = [
            "go", "build",
            "-buildmode=c-shared",
            "-o", output_path,
            "./ffi/analyzer.go"
        ]
        
        # Set environment variables
        env = os.environ.copy()
        env["CGO_ENABLED"] = "1"
        
        print(f"Building Go shared library: {' '.join(cmd)}")
        result = subprocess.run(
            cmd,
            cwd=go_src_dir,
            env=env,
            capture_output=True,
            text=True
        )
        
        if result.returncode != 0:
            raise RuntimeError(
                f"Failed to build Go shared library:\n"
                f"stdout: {result.stdout}\n"
                f"stderr: {result.stderr}"
            )
        
        print(f"Successfully built {output_name}")


# Read long description from README
readme_path = Path(__file__).parent / "README.md"
long_description = readme_path.read_text(encoding="utf-8")

go_source_dir = Path(__file__).parent.parent.absolute()
print(f"go_source_dir: {go_source_dir}")

setup(
    name="analyzer",
    version="0.1.0",
    author="saint1991",
    description="Python bindings for Analyzer",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/saint1991/antlr-editor",
    packages=find_packages(exclude=["examples", "tests"]),
    ext_modules=[
        GoExtension("analyzer.libanalyzer"),
    ],
    cmdclass={
        "build_ext": GoBuildExt,
    },
    python_requires=">=3.8",
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
        "Programming Language :: Go",
    ],
    keywords="antlr parser analyzer expression",
    project_urls={
        "Bug Reports": "https://github.com/saint1991/antlr-editor/issues",
        "Source": "https://github.com/saint1991/antlr-editor",
    },
)