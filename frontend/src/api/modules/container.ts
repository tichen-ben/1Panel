import http from '@/api';
import { ResPage, ReqPage } from '../interface';
import { Container } from '../interface/container';

export const getContainerPage = (params: Container.ContainerSearch) => {
    return http.post<ResPage<Container.ContainerInfo>>(`/containers/search`, params);
};

export const getContainerLog = (params: Container.ContainerLogSearch) => {
    return http.post<string>(`/containers/log`, params);
};

export const ContainerOperator = (params: Container.ContainerOperate) => {
    return http.post(`/containers/operate`, params);
};

export const getContainerInspect = (containerID: string) => {
    return http.get<string>(`/containers/detail/${containerID}`);
};

// repo
export const getRepoPage = (params: ReqPage) => {
    return http.post<ResPage<Container.RepoInfo>>(`/containers/repo/search`, params);
};
export const repoCreate = (params: Container.RepoCreate) => {
    return http.post(`/containers/repo`, params);
};
export const repoUpdate = (params: Container.RepoUpdate) => {
    return http.put(`/containers/repo/${params.id}`, params);
};
export const deleteRepo = (params: { ids: number[] }) => {
    return http.post(`/containers/repo/del`, params);
};