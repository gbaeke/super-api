# Workflows

## release.yaml

The release workflow runs when `release` activity occurs in the repository. More spcifically, the workflow runs when a new release is `published`. To learn more about creating releases, please see https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository. 

The example below illustrates what release.yaml does:
- you create a new release and give it a tag like 1.10.1
- when you publish the release, the workflow will run
- the workflow has one job called `build` with several steps; each step uses an action to perform a specific task:
    - Check out: check out the code
    - Cache Docker layers: cache the docker layers to speed up the build
    - Docker Meta: extract metadata from Git and GitHub events for use with Docker
    - Set up QEMU: QEMU allows you to build Docker images for other architectures other than the host; it is not needed for this workflow and only added as an example; we only build a Linux container image here
    - Set up Docker Build: install and use buildx; will also use buildkit
    - Login to GHCR (GitHub Container Registry): login to the GitHub Container Registry with username ${{ github.repository_owner }} (gbaeke) and the GitHub token as the password
    - Build and push the Docker image: use ./Dockerfile and set the tags to the tags in the output of Docker meta; the tags are the GitHub sha and the ref which is the published tag that triggers the workflow

**Note:** the other actions use cosign to sign the image; see https://github.com/sigstore/cosign for more information

The result of the workflow is a container image in GitHub Container Registry. For example, creating release 1.0.5 in the repository created a new image with the following tags:
- latest
- sha-75e4134
- 1.0.5

The digest of the image is: sha256:a3a5649b7592125f797b7d3810ff9239f12e3ed549b570bb33cfecc4e7d6fdb0

Because recent releases use cosign, GHCR also contains a signature with a name based on the digest of the image. In this case: sha256-a3a5649b7592125f797b7d3810ff9239f12e3ed549b570bb33cfecc4e7d6fdb0.sig.

