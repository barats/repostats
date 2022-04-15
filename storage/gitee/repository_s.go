// Copyright (c) [2022] [巴拉迪维 BaratSemet]
// [ohUrlShortener] is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
// 				 http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package gitee

import (
	gitee_model "repostats/model/gitee"
	"repostats/storage"
)

func BulkSaveRepos(repos []gitee_model.Repository) error {
	query := `INSERT INTO gitee.repos (id, full_name, human_name, path,name, url, owner_id,assigner_id, description, 
		html_url, ssh_url,forked_repo,default_branch, forks_count, stargazers_count, watchers_count,license, pushed_at, created_at, updated_at)
	VALUES(:id, :full_name, :human_name, :path, :name, :url, :owner.id, :assigner.id, :description, :html_url, :ssh_url, :forked_repo,
		:default_branch, :forks_count, :stargazers_count, :watchers_count, :license, :pushed_at, :created_at, :updated_at)
	ON CONFLICT (id) DO UPDATE SET id=EXCLUDED.id,full_name=EXCLUDED.full_name,human_name=EXCLUDED.human_name,path=EXCLUDED.path,
		url=EXCLUDED.url,owner_id=EXCLUDED.owner_id,assigner_id=EXCLUDED.assigner_id, description=EXCLUDED.description,
		html_url=EXCLUDED.html_url,ssh_url=EXCLUDED.ssh_url,forked_repo=EXCLUDED.forked_repo, forks_count=EXCLUDED.forks_count,
		stargazers_count=EXCLUDED.stargazers_count, watchers_count=EXCLUDED.watchers_count,
		license=EXCLUDED.license,pushed_at=EXCLUDED.pushed_at,created_at=EXCLUDED.created_at,updated_at=EXCLUDED.updated_at`
	return storage.DbNamedExec(query, repos)
}
