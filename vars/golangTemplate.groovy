def call(String podLabel, code) { podTemplate(
    cloud: 'ibm k8s',
    namespace: 'jenkins',
    label: podLabel,
    containers: [
      containerTemplate(
        name: 'golang',
        image: 'golang:1.13.0-alpine',
        ttyEnabled: true,
        command: 'cat')
]) {
code() }
}