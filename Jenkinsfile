#!groovy

lswci([node: 'docker']) {

    stage("Deploy") {
        if (env.BRANCH_NAME == 'master') {
            sshagent(['jenkins-ci-key']) {
                sh 'git remote add github git@github.com:Leaseweb/leaseweb-go-sdk-tmp.git'
                sh 'git branch -M master'
                sh 'git push -f github master'
            }
        }
    }
}