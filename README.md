# Panopticon

> The Panopticon is a type of institutional building and a system of control designed by the English philosopher and social theorist Jeremy Bentham in the late 18th century. The scheme of the design is to allow all inmates of an institution to be observed by a single watchman without the inmates being able to tell whether or not they are being watched.

## Usage
The repository has been setup in such a way that any user of AWS should be able to populate their own credentials into the project and have it deploy within their own environment. I will include `bash aliases` within the repository so that you may source your own variables and have them populate the terraform files upon run.

In order to run the code, you must find your way into the `scripts/` directory and make edits to the values of the `export`'s in there, at the time of this writing, the values to be changed are:

* `TF_VARS_account_name`
* `TF_VARS_account_email`
