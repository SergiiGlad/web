def call(String podLabel, code) { podTemplate(
    cloud: 'kubernetes-dev',
    namespace: 'jenkins',
    label: podLabel,
    containers: [
      containerTemplate(
        name: 'dotnet',
        image: 'microsoft/dotnet',
        ttyEnabled: true,
        command: 'cat')
]) {
code() }
}