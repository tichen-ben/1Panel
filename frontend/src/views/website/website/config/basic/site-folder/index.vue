<template>
    <div v-loading="loading">
        <div class="site-form-wrapper">
            <el-form class="site-form moblie-form" ref="siteForm" :model="update" label-width="100px">
                <el-form-item :label="$t('website.siteAlias')">
                    {{ website.alias }}
                </el-form-item>
                <el-form-item :label="$t('website.primaryPath')">
                    <el-space wrap>
                        {{ website.sitePath }}
                        <el-button type="primary" link @click="toFolder(website.sitePath)">
                            <el-icon>
                                <FolderOpened />
                            </el-icon>
                        </el-button>
                    </el-space>
                </el-form-item>
                <el-form-item v-if="configDir" :label="$t('website.runDir')">
                    <el-space wrap>
                        <el-select v-model="update.siteDir">
                            <el-option :label="'/'" :value="'/'"></el-option>
                            <el-option
                                v-for="(item, index) in dirs"
                                :label="item"
                                :value="item"
                                :key="index"
                            ></el-option>
                        </el-select>
                        <el-button type="primary" @click="submit(siteForm)">
                            {{ $t('nginx.saveAndReload') }}
                        </el-button>
                    </el-space>
                </el-form-item>
                <el-form-item v-if="configDir" :label="$t('website.userGroup')">
                    <el-space wrap>
                        <el-input v-model="updatePermission.user" class="user-num-input">
                            <template #prepend>{{ $t('website.user') }}</template>
                        </el-input>
                        <el-input v-model="updatePermission.group" class="user-num-input">
                            <template #prepend>{{ $t('website.uGroup') }}</template>
                        </el-input>
                        <el-button type="primary" @click="submitPermission()">
                            {{ $t('commons.button.save') }}
                        </el-button>
                    </el-space>
                </el-form-item>
            </el-form>
            <el-alert :closable="false" v-if="configDir">
                <template #default>
                    <span class="warnHelper">{{ $t('website.runDirHelper') }}</span>
                    <span class="warnHelper">{{ $t('website.runUserHelper') }}</span>
                </template>
            </el-alert>
            <br />
            <el-descriptions :title="$t('website.folderTitle')" :column="1" border>
                <el-descriptions-item label="waf">{{ $t('website.wafFolder') }}</el-descriptions-item>
                <el-descriptions-item label="ssl">{{ $t('website.sslFolder') }}</el-descriptions-item>
                <el-descriptions-item label="log">{{ $t('website.logFolder') }}</el-descriptions-item>
                <el-descriptions-item label="index">{{ $t('website.indexFolder') }}</el-descriptions-item>
            </el-descriptions>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { GetFilesList } from '@/api/modules/files';
import { GetWebsite, UpdateWebsiteDir, UpdateWebsiteDirPermission } from '@/api/modules/website';
import i18n from '@/lang';
import { MsgSuccess } from '@/utils/message';
import { FormInstance } from 'element-plus';
import { computed, onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter();

const props = defineProps({
    id: {
        type: Number,
        default: 0,
    },
});
const websiteId = computed(() => {
    return Number(props.id);
});
const website = ref<any>({});
const loading = ref(false);
const configDir = ref(false);
const update = reactive({
    id: 0,
    siteDir: '/',
});
const updatePermission = reactive({
    id: 0,
    user: '1000',
    group: '1000',
});
const siteForm = ref<FormInstance>();
const dirReq = reactive({
    path: '/',
    expand: true,
    showHidden: false,
    page: 1,
    pageSize: 100,
    search: '',
    containSub: false,
    dir: true,
});
const dirs = ref([]);

const search = () => {
    loading.value = true;
    GetWebsite(websiteId.value)
        .then((res) => {
            website.value = res.data;
            update.id = website.value.id;
            update.siteDir = website.value.siteDir;
            updatePermission.id = website.value.id;
            updatePermission.group = website.value.group === '' ? '1000' : website.value.group;
            updatePermission.user = website.value.user === '' ? '1000' : website.value.user;
            if (website.value.type === 'static' || website.value.runtimeID > 0) {
                configDir.value = true;
                dirReq.path = website.value.sitePath + '/index';
                getDirs();
            }
        })
        .finally(() => {
            loading.value = false;
        });
};

const submit = async (formEl: FormInstance | undefined) => {
    if (!formEl) return;
    await formEl.validate((valid) => {
        if (!valid) {
            return;
        }
        loading.value = true;
        UpdateWebsiteDir(update)
            .then(() => {
                MsgSuccess(i18n.global.t('commons.msg.updateSuccess'));
                search();
            })
            .finally(() => {
                loading.value = false;
            });
    });
};

const submitPermission = async () => {
    if (updatePermission.user === '' || updatePermission.group === '') {
        return;
    }
    loading.value = true;
    UpdateWebsiteDirPermission(updatePermission)
        .then(() => {
            MsgSuccess(i18n.global.t('commons.msg.updateSuccess'));
            search();
        })
        .finally(() => {
            loading.value = false;
        });
};

const getDirs = async () => {
    loading.value = true;
    await GetFilesList(dirReq)
        .then((res) => {
            dirs.value = [];
            const items = res.data.items || [];
            for (const item of items) {
                dirs.value.push(item.name);
            }
        })
        .finally(() => {
            loading.value = false;
        });
};

const initData = () => {
    dirs.value = [];
};

const toFolder = (folder: string) => {
    router.push({ path: '/hosts/files', query: { path: folder } });
};

onMounted(() => {
    initData();
    search();
});
</script>

<style lang="scss" scoped>
.site-form-wrapper {
    min-width: 600px;
    width: 60%;
    padding: 20px;
}
.site-form {
    :deep(.el-form-item__label) {
        padding-right: 20px !important;
        box-sizing: content-box;
    }
    .user-num-input {
        width: 190px;
    }
}
.warnHelper {
    white-space: pre-line;
    display: block;
}
</style>
