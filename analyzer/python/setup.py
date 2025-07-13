#!/usr/bin/env python3
"""
Setup script for the ANTLR Expression Analyzer Python bindings
"""

from setuptools import setup, find_packages
from pathlib import Path

# Read the README file
readme_path = Path(__file__).parent.parent / "README.md"
long_description = ""
if readme_path.exists():
    long_description = readme_path.read_text(encoding="utf-8")

setup(
    name="antlr-expression-analyzer",
    version="1.0.0",
    description="Python bindings for the ANTLR Expression Analyzer",
    long_description=long_description,
    long_description_content_type="text/markdown",
    author="ANTLR Editor Project",
    python_requires=">=3.7",
    packages=find_packages(),
    include_package_data=True,
    package_data={
        "": ["*.so", "*.dylib", "*.dll", "*.h"],
    },
    classifiers=[
        "Development Status :: 4 - Beta",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
        "Topic :: Software Development :: Compilers",
        "Topic :: Text Processing :: Linguistic",
    ],
    keywords="antlr parser expression analyzer syntax highlighting",
    project_urls={
        "Source": "https://github.com/saint1991/antlr-editor",
        "Bug Reports": "https://github.com/saint1991/antlr-editor/issues",
    },
)