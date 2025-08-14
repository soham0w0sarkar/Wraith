<script>
	import { createEventDispatcher, onMount } from 'svelte';

	let {
		item,
		iconTheme = $bindable(null),
		expanded = false,
		level = 0,
		isRoot = false
	} = $props();

	const dispatch = createEventDispatcher();

	onMount(async () => {
		if (!iconTheme) {
			try {
				const response = await fetch('/mocha/theme.json');
				iconTheme = await response.json();
			} catch (error) {
				console.error('Failed to load icon theme:', error);
			}
		}
	});

	function resolveIcon(item, iconTheme, expanded = false, isRoot = false) {
		if (!iconTheme?.iconDefinitions) return null;

		const { name, type, extension, languageId } = item;
		let iconKey = null;

		if (type === 'folder') {
			if (isRoot) {
				iconKey = expanded ? iconTheme.rootFolderExpanded : iconTheme.rootFolder;
			} else {
				if (expanded && iconTheme.folderNamesExpanded?.[name]) {
					iconKey = iconTheme.folderNamesExpanded[name];
				} else if (iconTheme.folderNames?.[name]) {
					iconKey = iconTheme.folderNames[name];
				} else {
					iconKey = expanded ? iconTheme.folderExpanded : iconTheme.folder;
				}
			}
		} else {
			if (iconTheme.fileNames?.[name]) {
				iconKey = iconTheme.fileNames[name];
			}
			else if (extension && iconTheme.fileExtensions?.[extension]) {
				iconKey = iconTheme.fileExtensions[extension];
			}
			else if (languageId && iconTheme.languageIds?.[languageId]) {
				iconKey = iconTheme.languageIds[languageId];
			}
			else {
				iconKey = iconTheme.file;
			}
		}

		return iconKey ? iconTheme.iconDefinitions[iconKey] : null;
	}

	const icon = $derived(resolveIcon(item, iconTheme, expanded, isRoot));
	const iconPath = $derived(icon?.iconPath);

	function handleClick() {
		if (item.type === 'folder') {
			const newExpanded = !expanded;
			dispatch('toggle', { item, expanded: newExpanded });
		} else {
			dispatch('select', { item });
		}
	}

	function handleKeydown(event) {
		if (event.key === 'Enter' || event.key === ' ') {
			event.preventDefault();
			handleClick();
		}
	}
</script>

<div
	class="flex items-center gap-1.5 px-2 py-1 font-mono text-sm text-gray-200 cursor-pointer select-none transition-colors border-l-2 border-transparent min-h-6 hover:bg-white/10 focus:outline-none focus:border-l-blue-400 focus:bg-blue-400/15"
	class:folder={item.type === 'folder'}
	class:file={item.type === 'file'}
	class:expanded
	style="padding-left: {level * 16 + 8}px"
	onclick={handleClick}
	onkeydown={handleKeydown}
	tabindex="0"
	role="treeitem"
    aria-selected="true"
	aria-expanded={item.type === 'folder' ? expanded : undefined}
>
	{#if item.type === 'folder'}
		<div class="w-[9px] h-[9px] flex items-center justify-center transition-transform text-gray-400 flex-shrink-0" class:rotate-90={expanded}>
			<img src="/mocha/icons/_down_arrow.svg" alt="Toggle folder" class="w-[9px] h-[9px]" />
		</div>
	{:else}
		<div class="w-3 h-3 flex-shrink-0"></div>
	{/if}

	<div class="w-4 h-4 flex items-center justify-center text-gray-300 flex-shrink-0">
		{#if iconPath}
			<img src="/mocha/{iconPath}" alt="" class="w-4 h-4" />
		{:else}
			{#if item.type === 'folder'}
				<svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
					<path
						d="M14.5 3H7.71l-.85-.85A.5.5 0 0 0 6.5 2h-5a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5v-10a.5.5 0 0 0-.5-.5Z"
					/>
				</svg>
			{:else}
				<svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
					<path d="M4 0h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2Z" />
				</svg>
			{/if}
		{/if}
	</div>

	<!-- File/folder name -->
	<span class="flex-1 overflow-hidden text-ellipsis whitespace-nowrap">{item.name}</span>
</div>