#include <Python.h>

// ===== START PYTHON PART =====

/* Will come from go */
PyObject* command(PyObject* , PyObject*);

/*
To shim go's missing variadic function support.
Ref https://docs.python.org/3/c-api/arg.html
Reminder from docs about memory management when parsing arguments:
> In general, when a format sets a pointer to a buffer,
> the buffer is managed by the corresponding Python object,
> and the buffer shares the lifetime of this object. You
> won’t have to release any memory yourself. The only
> exceptions are es, es#, et and et#.
*/
int PyArg_ParseTuple_ss(PyObject* args, char** a, char** b) {
    return PyArg_ParseTuple(args, "ss", a, b);
}

static struct PyMethodDef methods[] = {
    {
    "command",
    (PyCFunction)command,
    METH_VARARGS,
    "Runs a command."
    },
    {NULL, NULL, 0, NULL}
};

static struct PyModuleDef module = {
    PyModuleDef_HEAD_INIT,
    "_playground",
    "Go functions for Python",
    -1,
    methods
};

PyMODINIT_FUNC PyInit__playground(void) {
    PyObject *m;

    m = PyModule_Create(&module);
    if (m == NULL)
        return NULL;

    return m;
}
