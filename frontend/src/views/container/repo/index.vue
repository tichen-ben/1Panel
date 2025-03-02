<template>
    <div v-loading="loading">
        <el-card v-if="dockerStatus != 'Running'" class="mask-prompt">
            <span>{{ $t('container.serviceUnavailable') }}</span>
            <el-button type="primary" link class="bt" @click="goSetting">【 {{ $t('container.setting') }} 】</el-button>
            <span>{{ $t('container.startIn') }}</span>
        </el-card>

        <LayoutContent :title="$t('container.repo')" :class="{ mask: dockerStatus != 'Running' }">
            <template #toolbar>
                <el-row>
                    <el-col :span="16">
                        <el-button type="primary" @click="onOpenDialog('add')">
                            {{ $t('container.createRepo') }}
                        </el-button>
                    </el-col>
                    <el-col :span="8">
                        <TableSetting @search="search()" />
                        <div class="search-button">
                            <el-input
                                v-model="searchName"
                                clearable
                                @clear="search()"
                                suffix-icon="Search"
                                @keyup.enter="search()"
                                @change="search()"
                                :placeholder="$t('commons.button.search')"
                            ></el-input>
                        </div>
                    </el-col>
                </el-row>
            </template>
            <template #main>
                <ComplexTable
                    :pagination-config="paginationConfig"
                    v-model:selects="selects"
                    :data="data"
                    @search="search"
                >
                    <el-table-column :label="$t('commons.table.name')" prop="name" min-width="60" />
                    <el-table-column
                        :label="$t('container.downloadUrl')"
                        show-overflow-tooltip
                        prop="downloadUrl"
                        min-width="100"
                        fix
                    />
                    <el-table-column :label="$t('container.protocol')" prop="protocol" min-width="60" fix />
                    <el-table-column :label="$t('commons.table.status')" prop="status" min-width="60" fix>
                        <template #default="{ row }">
                            <el-tag v-if="row.status === 'Success'" type="success">
                                {{ $t('commons.status.success') }}
                            </el-tag>
                            <el-tooltip v-else effect="dark" placement="bottom">
                                <template #content>
                                    <div style="width: 300px; word-break: break-all">{{ row.message }}</div>
                                </template>
                                <el-tag type="danger">{{ $t('commons.status.failed') }}</el-tag>
                            </el-tooltip>
                        </template>
                    </el-table-column>
                    <el-table-column
                        prop="createdAt"
                        :label="$t('commons.table.createdAt')"
                        min-width="80"
                        fix
                        :formatter="dateFormat"
                    />
                    <fu-table-operations :buttons="buttons" :label="$t('commons.table.operate')" />
                </ComplexTable>
            </template>
        </LayoutContent>
        <OperatorDialog @search="search" ref="dialogRef" />
    </div>
</template>

<script lang="ts" setup>
import TableSetting from '@/components/table-setting/index.vue';
import OperatorDialog from '@/views/container/repo/operator/index.vue';
import { reactive, onMounted, ref } from 'vue';
import { dateFormat } from '@/utils/util';
import { Container } from '@/api/interface/container';
import { checkRepoStatus, deleteImageRepo, loadDockerStatus, searchImageRepo } from '@/api/modules/container';
import i18n from '@/lang';
import router from '@/routers';
import { ElMessageBox } from 'element-plus';

const loading = ref();
const data = ref();
const selects = ref<any>([]);
const paginationConfig = reactive({
    currentPage: 1,
    pageSize: 10,
    total: 0,
});
const searchName = ref();

const dockerStatus = ref('Running');
const loadStatus = async () => {
    loading.value = true;
    await loadDockerStatus()
        .then((res) => {
            loading.value = false;
            dockerStatus.value = res.data;
            if (dockerStatus.value === 'Running') {
                search();
            }
        })
        .catch(() => {
            dockerStatus.value = 'Failed';
            loading.value = false;
        });
};
const goSetting = async () => {
    router.push({ name: 'ContainerSetting' });
};

const search = async () => {
    let params = {
        info: searchName.value,
        page: paginationConfig.currentPage,
        pageSize: paginationConfig.pageSize,
    };
    loading.value = true;
    await searchImageRepo(params)
        .then((res) => {
            loading.value = false;
            data.value = res.data.items || [];
            paginationConfig.total = res.data.total;
        })
        .catch(() => {
            loading.value = false;
        });
};
const dialogRef = ref();
const onOpenDialog = async (
    title: string,
    rowData: Partial<Container.RepoInfo> = {
        auth: true,
        protocol: 'http',
    },
) => {
    let params = {
        title,
        rowData: { ...rowData },
    };
    dialogRef.value!.acceptParams(params);
};

const onDelete = async (row: Container.RepoInfo) => {
    ElMessageBox.confirm(i18n.global.t('commons.msg.delete'), i18n.global.t('commons.button.delete'), {
        confirmButtonText: i18n.global.t('commons.button.confirm'),
        cancelButtonText: i18n.global.t('commons.button.cancel'),
    }).then(async () => {
        await deleteImageRepo({ ids: [row.id] });
        search();
    });
};

const onCheckConn = async (row: Container.RepoInfo) => {
    loading.value = true;
    await checkRepoStatus(row.id)
        .then(() => {
            loading.value = false;
            search();
        })
        .catch(() => {
            loading.value = false;
        });
};

const buttons = [
    {
        label: i18n.global.t('commons.button.sync'),
        disabled: (row: Container.RepoInfo) => {
            return row.downloadUrl === 'docker.io';
        },
        click: (row: Container.RepoInfo) => {
            onCheckConn(row);
        },
    },
    {
        label: i18n.global.t('commons.button.edit'),
        disabled: (row: Container.RepoInfo) => {
            return row.downloadUrl === 'docker.io';
        },
        click: (row: Container.RepoInfo) => {
            onOpenDialog('edit', row);
        },
    },
    {
        label: i18n.global.t('commons.button.delete'),
        disabled: (row: Container.RepoInfo) => {
            return row.downloadUrl === 'docker.io';
        },
        click: (row: Container.RepoInfo) => {
            onDelete(row);
        },
    },
];

onMounted(() => {
    loadStatus();
});
</script>
