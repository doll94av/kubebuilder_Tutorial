# project

Cool tutorial overall, helped a lot in understanding controllers and how API groups are developed in some capacity. I verified that my controller was launching and managing my CRD instances, there are multiple API versions available (v1/v2) and you can feed in a custom config file (config/manager/controller_manager_config.yaml)


# modifications

There were a few issues I ran into when running through the tutorial. The first being that the tutorial was referncing functions in the log package that were not available, so I commented out any log lines that Make was throwing errors for. The other issue was that if you go through with the multiple API groups, you run into a situation where the default Makefile is using 'kubectl apply -f' and throwing an 'annotation to long' type of error, this can be worked around by modifying the Makefiles 'install' command to use kubectl create. 
