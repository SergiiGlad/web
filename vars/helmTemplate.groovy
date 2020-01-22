def call(String podLabel, code) { podTemplate(
    cloud: 'kubernetes,
    namespace: 'jenkins',
    label: podLabel,
    containers: [
      containerTemplate(
        name: 'helm',
        image: 'lachlanevenson/k8s-helm:v2.16.1',
        ttyEnabled: true,
        command: 'cat'),  
      containerTemplate(
        name: 'kubectl',
        image: 'lachlanevenson/k8s-kubectl:v1.16.4',
        ttyEnabled: true,
        command: 'cat')  
]) {
code() }
}
