#!groovy

lswci([node: 'docker']) {
    stage("Deploy") {
        sshagent(['jenkins-ci-key']) {
            sh 'git remote add github git@github.com:Leaseweb/leaseweb-go-sdk-tmp.git'
            sh 'git branch -M master'
            sh 'git push -u github master'
        }
    }
}