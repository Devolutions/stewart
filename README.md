Terraform wrapper to simply the use of workspace/environment variables files setups.

The wrapper was kept very simple, it doesn't really validate any inputs. It relies on Terraform
to do that.

# Installation

+ Grab the binary for your OS/arch
+ Drop it somewhere that's in your $PATH

# Usage

Stewart requires 2 arguments:

+ a Terraform subcommand
+ an environment

It's also possible to pass optional arguments like the example below to import a resource for example:

~~~
stewart import development azurerm_resource_group.main[\"hello\"] /subscriptions/$some_ID/resourceGroups/hello-world-development
~~~

Basically, an subcommand is passed straight to Terraform with no validation. Terraform will raise an error if the command
doesn't exist. As for the environment, it will try to use the file `environments/$environment.tfvars`, so if that doesn't
exist, Terraform will also raise an error.

Lastly, it automatically applies configurations (-auto-approve).

# Stewart? What kind of name is that?

~~~
Author Jack Williamson is credited with inventing and popularizing the term "terraform".
In July 1942, under the pseudonym Will Stewart, Williamson published a science fiction
novella entitled "Collision Orbit" in Astounding Science-Fiction magazine.
~~~

So the wrapper was named after the first person to come up with the terraforming concept.
