# project

Cool tutorial overall, helped a lot in understanding controllers and how API groups are developed in some capacity. I verified that my controller was launching:

austin@ubuntu:~/project$ make run ENABLE_WEBHOOKS=false
GOBIN=/home/austin/project/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0
/home/austin/project/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
/home/austin/project/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
main.go
go vet ./...
go run ./main.go
I0606 11:50:20.836754   95832 request.go:665] Waited for 1.040817256s due to client-side throttling, not priority and fairness, request: GET:https://127.0.0.1:38141/apis/infrastructure.cluster.x-k8s.io/v1beta1?timeout=32s
1.6545306208901632e+09  INFO    controller-runtime.metrics      Metrics server is starting to listen    {"addr": ":8080"}
1.6545306208905056e+09  INFO    setup   starting manager
1.6545306208923419e+09  INFO    Starting server {"kind": "health probe", "addr": "[::]:8081"}
1.6545306208923464e+09  INFO    Starting server {"path": "/metrics", "kind": "metrics", "addr": "[::]:8080"}
1.6545306209929543e+09  INFO    controller.cronjob      Starting EventSource    {"reconciler group": "batch.tutorial.kubebuilder.io", "reconciler kind": "CronJob", "source": "kind source: *v1.CronJob"}
1.6545306209930873e+09  INFO    controller.cronjob      Starting EventSource    {"reconciler group": "batch.tutorial.kubebuilder.io", "reconciler kind": "CronJob", "source": "kind source: *v1.Job"}
1.6545306209930909e+09  INFO    controller.cronjob      Starting Controller     {"reconciler group": "batch.tutorial.kubebuilder.io", "reconciler kind": "CronJob"}
1.6545306210960605e+09  INFO    controller.cronjob      Starting workers        {"reconciler group": "batch.tutorial.kubebuilder.io", "reconciler kind": "CronJob", "worker count": 1}


there are multiple API versions available (v1/v2) and main has been updated to accept acustom config file (config/manager/controller_manager_config.yaml)


# modifications

There were a few issues I ran into when running through the tutorial. The first being that the tutorial was referncing functions in the log package that were not available, so I commented out any log lines that Make was throwing errors for. The other issue was that if you go through with the multiple API groups, you run into a situation where the default Makefile is using 'kubectl apply -f' and throwing an 'annotation to long' type of error, this can be worked around by modifying the Makefiles 'install' command to use kubectl create. 
