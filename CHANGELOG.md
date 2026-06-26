# Changelog

## [0.2.1](https://github.com/NitriKx/provider-cloudinary/compare/v0.2.0...v0.2.1) (2026-06-26)


### Bug Fixes

* **ci:** [INFRA-6551] embed runtime image in xpkg via crossplane CLI ([40481fc](https://github.com/NitriKx/provider-cloudinary/commit/40481fca69af76aa396f75735ca45129c05035f4))

## [0.2.0](https://github.com/NitriKx/provider-cloudinary/compare/v0.1.0...v0.2.0) (2026-06-26)


### Features

* add cloudinary_upload_preset managed resource ([0087dbb](https://github.com/NitriKx/provider-cloudinary/commit/0087dbb23681722dbb41eb5db4b1cc4a317819ca))
* add cloudinary_upload_preset managed resource ([c5abffa](https://github.com/NitriKx/provider-cloudinary/commit/c5abffa2a2c7c65a30bc5dd15213b0b3e6b77741))


### Bug Fixes

* **ci:** [INFRA-6551] add required packages block to release-please config ([2533358](https://github.com/NitriKx/provider-cloudinary/commit/2533358c84b5d3efaaeb06649726b0e21c53eaab))
* **ci:** [INFRA-6551] copy terraformrc.hcl into build context before docker build ([0ba0620](https://github.com/NitriKx/provider-cloudinary/commit/0ba0620f9e6def080d0f3fd3e3924190834493f3))
* **ci:** [INFRA-6551] embed controller image ref in xpkg build ([12ef135](https://github.com/NitriKx/provider-cloudinary/commit/12ef13579ef65fc2d7dd77699ccc633b841ef4b2))
* **ci:** [INFRA-6551] fix release-please manifest component key ([350c425](https://github.com/NitriKx/provider-cloudinary/commit/350c4255321e4d1a289176b4379ad1d9919b1941))
* **ci:** [INFRA-6551] use RP_GITHUB_TOKEN for release-please PR creation ([15280ff](https://github.com/NitriKx/provider-cloudinary/commit/15280ffa675d4fa04d4ab0d6fc322851e4f018b6))
* **ci:** correct terraformrc.hcl path in provider image build ([#2](https://github.com/NitriKx/provider-cloudinary/issues/2)) ([e39e213](https://github.com/NitriKx/provider-cloudinary/commit/e39e2138b082b5814fdae3fbb1595e8495e95d02))
* **ci:** fix up xpkg push syntax (tag is positional arg, file is -f flag) ([028e887](https://github.com/NitriKx/provider-cloudinary/commit/028e887055df446f438360cd89860227fcadd035))
* **ci:** lowercase provider name in provider-metadata.yaml ([d1f2f28](https://github.com/NitriKx/provider-cloudinary/commit/d1f2f283aad5112e07335107cd1e3794d2a8afcc))
* **ci:** replace crossplane-contrib publish workflow with direct GHCR push ([910f7f1](https://github.com/NitriKx/provider-cloudinary/commit/910f7f140b8b4af0bc887c44daa757b25769df3d))
* **ci:** rewrite publish workflow with explicit build/push steps ([6b74546](https://github.com/NitriKx/provider-cloudinary/commit/6b74546e6456b4a26eb101c44949adf316cebf83))
* use NoOpProviderScheduler and lowercase provider source address ([d17fe6f](https://github.com/NitriKx/provider-cloudinary/commit/d17fe6f64d6a4543d25e1058be3b2ff9d48f17ec))
