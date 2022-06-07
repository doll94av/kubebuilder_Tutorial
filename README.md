# project

Cool tutorial overall, helped a lot in understanding controllers and how API groups are developed in some capacity. 


# modifications

There were a few issues I ran into when running through the tutorial. The first being that the tutorial was referncing functions in the log package that were not available, so I commented out any log lines that Make was throwing errors for. The other issue was that if you go through with the multiple API groups, you run into a situation where the default Makefile is using 'kubectl apply -f' and throwing an 'annotation to long' type of error, this can be worked around by modifying the Makefiles 'install' command to use kubectl create. 
