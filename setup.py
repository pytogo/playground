"""The setup script."""

from setuptools import setup, Extension

requirements = []

test_requirements = ["pytest>=3", ]

setup(
    author="Sebastian Ziemann",
    author_email="corka149@mailbox.org",
    python_requires=">=3.6",
    description="Trial and test ground",
    install_requires=requirements,
    license="MIT license",
    include_package_data=True,
    keywords="playground",
    name="playground",
    py_modules=["playground"],
    test_suite="tests",
    tests_require=test_requirements,
    url="https://github.com/pytogo/playground",
    version="0.0.1",
    zip_safe=False,
    # Go part
    setup_requires=['setuptools-golang'],
    build_golang={'root': 'github.com/pytogo/playground'},
    ext_modules=[
        Extension(
            "_playground", ["main.go"],
            py_limited_api=True, define_macros=[('Py_LIMITED_API', None)],
        )
    ]
)
