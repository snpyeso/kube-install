package kilib

import (
    "os"
)


// Generate orchestration for install.
func InstallYML(mode string, softDir string, currentDir string, currentUser string, logName string, upgradeKernel string, osTypeResult string, k8sVer string, cniPlugin string) {
    var osCompatibilityLayout,runtime,cniPluginStr string
    if upgradeKernel  == "yes" {
        if osTypeResult == "centos7" || osTypeResult == "rhel7" {
            osCompatibilityLayout = "- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/kernel\n"
        }
    }
    if k8sVer == "1.24" {
        runtime = "runtime"
    } else {
        runtime = "docker"
    }
    if cniPlugin == "" {
        cniPluginStr = "- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000network/flannel\n"
    } else {
        cniPluginStr = "- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000network/"+cniPlugin+"\n"
    }
    install_file, err := os.Create(softDir+"/k8scluster-install.yml")
    CheckErr(err,currentDir,logName,mode)
    defer install_file.Close()
    install_file.WriteString("- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/genfile\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/all\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/"+runtime+"\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x000certificate/copycfssl\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x000certificate/createssl\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x000certificate/fetchssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000storage\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/apiserver\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/api-rbac\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/scheduler\n"+cniPluginStr+"- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000node/kubelet\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000node/approve-csr\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000node/kube-proxy\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000addons\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/admintoken\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000action/pushimages\n"+osCompatibilityLayout+"- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000finish/install\n")
}

// Generate orchestration for install one master.
func OnemasterInstallYML(mode string, softDir string, currentDir string, currentUser string, logName string, upgradeKernel string, osTypeResult string, k8sVer string, cniPlugin string) {
    var osCompatibilityLayout,runtime,cniPluginStr string
    if upgradeKernel  == "yes" { 
        if osTypeResult == "centos7" || osTypeResult == "rhel7" {
            osCompatibilityLayout = "- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/kernel\n"
        }
    }
    if k8sVer == "1.24" {
        runtime = "runtime"
    } else {
        runtime = "docker"
    }
    if cniPlugin == "" {
        cniPluginStr = "- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000network/flannel\n"
    } else {
        cniPluginStr = "- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000network/"+cniPlugin+"\n"
    }
    onemasterinstall_file, err := os.Create(softDir+"/k8scluster-onemasterinstall.yml")
    CheckErr(err,currentDir,logName,mode)
    defer onemasterinstall_file.Close()
    onemasterinstall_file.WriteString("- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/genfile\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/all\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/"+runtime+"\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x000certificate/copycfssl\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x000certificate/createssl\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x000certificate/fetchssl\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000storage\n- remote_user: root\n  hosts: master,node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/apiserver\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/api-rbac\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/scheduler\n"+cniPluginStr+"- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000node/kubelet\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000node/approve-csr\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000node/kube-proxy\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000addons\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/admintoken\n- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000action/pushimages\n"+osCompatibilityLayout+"- remote_user: root\n  hosts: node\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000finish/install\n")
}

// Generate orchestration for add node.
func AddnodeYML(mode string, softDir string, currentDir string, currentUser string, logName string, upgradeKernel string, osTypeResult string, k8sVer string) {
    var osCompatibilityLayout,runtime string
    if upgradeKernel  == "yes" { 
        if osTypeResult == "centos7" || osTypeResult == "rhel7" {
            osCompatibilityLayout = "- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/kernel\n"
        }
    }
    if k8sVer == "1.24" {
        runtime = "runtime"
    } else {
        runtime = "docker"
    }
    addnode_file, err := os.Create(softDir+"/k8scluster-addnode.yml")
    CheckErr(err,currentDir,logName,mode)
    defer addnode_file.Close()
    addnode_file.WriteString("- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/genfile\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/all\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/"+runtime+"\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000node/kubelet\n- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000node/kube-proxy\n- remote_user: root\n  hosts: master1\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000node/approve-csr\n"+osCompatibilityLayout+"- remote_user: root\n  hosts: addnode\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000finish/addnode\n")
}

// Generate orchestration for delete node.
func DelnodeYML(mode string, softDir string, currentDir string, currentUser string, logName string, force bool) {
    var clearFileStr string
    if force{
        clearFileStr = "- remote_user: root\n  hosts: delnode\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000action/cleargarbage\n"
    }
    delnode_file, err := os.Create(softDir+"/k8scluster-delnode.yml")
    CheckErr(err,currentDir,logName,mode)
    defer delnode_file.Close()
    delnode_file.WriteString("- remote_user: root\n  hosts: delnode\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000action/delnode\n"+clearFileStr)
}

// Generate orchestration for rebuild master.
func RebuildmasterYML(mode string, softDir string, currentDir string, currentUser string, logName string) {
    rebuildmaster_file, err := os.Create(softDir+"/k8scluster-rebuildmaster.yml")
    CheckErr(err,currentDir,logName,mode)
    defer rebuildmaster_file.Close()
    rebuildmaster_file.WriteString("- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/genfile\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000000base/all\n- remote_user: root\n  hosts: etcd\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x0000000storage\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/kubectl\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/apiserver\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/controller-manager\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/scheduler\n- remote_user: root\n  hosts: master\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000master/admintoken\n")
}

// Generate orchestration for delete master.
func DelmasterYML(mode string, softDir string, currentDir string, currentUser string, logName string, force bool) {
    var clearFileStr string
    if force{
        clearFileStr = "- remote_user: root\n  hosts: delmaster\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000action/cleargarbage\n"
    }
    delmaster_file, err := os.Create(softDir+"/k8scluster-delmaster.yml")
    CheckErr(err,currentDir,logName,mode)
    defer delmaster_file.Close()
    delmaster_file.WriteString("- remote_user: root\n  hosts: delmaster\n  gather_facts: no\n  roles:\n    - "+softDir+"/sys/0x00000000action/delmaster\n"+clearFileStr)
}


