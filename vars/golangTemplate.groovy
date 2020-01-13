def call(String podLabel, code) { podTemplate(
    cloud: 'ibm kubernetes',
    namespace: 'jenkins',
    label: podLabel,
    containers: [
      containerTemplate(
        name: golang,
        image: golang:1.13.0-alpine,
        command: 'cat',
        ttyEnabled: true,
        name: docker-dind
        image: docker:stable-dind
        securityContext:
          privileged: true
        name: helm
        image: lachlanevenson/k8s-helm:v2.16.1
        command: 'cat',
        ttyEnabled: true,
        name: kubectl
        image: lachlanevenson/k8s-kubectl:v1.16.4
        command: 'cat',
        ttyEnabled: true)
]) {
code() }
}