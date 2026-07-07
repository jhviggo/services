// Make sure github dependencies are installed
const util = require('util');
const exec = util.promisify(require('child_process').exec);
exec('npm install @actions/core@1.9.1 @actions/github@5.1.1').then(() => {
    const core = require('@actions/core');
    const github = require('@actions/github');
    
    try {
        const inputStr = core.getInput('branch');
        const lowercase = inputStr.toLowerCase();
        core.exportVariable("BRANCH", lowercase);
        core.exportVariable("GITHUB_SHA_HASH", github.context.sha);
    } catch (error) {
        core.setFailed(error.message);
    }
});