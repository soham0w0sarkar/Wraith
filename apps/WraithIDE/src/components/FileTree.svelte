<script>
	import FileTreeItem from './FileTreeItem.svelte';

	const fileTree = [
		{
			name: 'src',
			type: 'folder',
			children: [
                {
                    name: 'docker',
                    type: 'folder',
                    children: [
                        {
                            name: 'Dockerfile',
                            type: 'file',
                            extension: 'dockerfile',
                            languageId: 'dockerfile'
                        }]
                },
				{
					name: 'components',
					type: 'folder',
					children: [
						{
							name: 'FileTreeItem.svelte',
							type: 'file',
							extension: 'svelte',
							languageId: 'svelte'
						},
						{
							name: 'Button.svelte',
							type: 'file',
							extension: 'svelte',
							languageId: 'svelte'
						},
						{
							name: 'Modal.svelte',
							type: 'file',
							extension: 'svelte',
							languageId: 'svelte'
						}
					]
				},
				{
					name: 'routes',
					type: 'folder',
					children: [
						{
							name: '+layout.svelte',
							type: 'file',
							extension: 'svelte',
							languageId: 'svelte'
						},
						{
							name: '+page.svelte',
							type: 'file',
							extension: 'svelte',
							languageId: 'svelte'
						},
						{
							name: 'about',
							type: 'folder',
							children: [
								{
									name: '+page.svelte',
									type: 'file',
									extension: 'svelte',
									languageId: 'svelte'
								}
							]
						}
					]
				},
				{
					name: 'lib',
					type: 'folder',
					children: [
						{
							name: 'utils.js',
							type: 'file',
							extension: 'js',
							languageId: 'javascript'
						},
						{
							name: 'api.js',
							type: 'file',
							extension: 'js',
							languageId: 'javascript'
						}
					]
				},
				{
					name: 'app.html',
					type: 'file',
					extension: 'html',
					languageId: 'html'
				},
				{
					name: 'app.css',
					type: 'file',
					extension: 'css',
					languageId: 'css'
				}
			]
		},
		{
			name: 'static',
			type: 'folder',
			children: [
				{
					name: 'images',
					type: 'folder',
					children: [
						{
							name: 'logo.png',
							type: 'file',
							extension: 'png',
							languageId: 'image'
						},
						{
							name: 'banner.jpg',
							type: 'file',
							extension: 'jpg',
							languageId: 'image'
						}
					]
				},
				{
					name: 'favicon.ico',
					type: 'file',
					extension: 'ico',
					languageId: 'image'
				}
			]
		},
		{
			name: 'package.json',
			type: 'file',
			extension: 'json',
			languageId: 'json'
		},
		{
			name: 'svelte.config.js',
			type: 'file',
			extension: 'js',
			languageId: 'javascript'
		},
		{
			name: 'vite.config.js',
			type: 'file',
			extension: 'js',
			languageId: 'javascript'
		},
		{
			name: 'README.md',
			type: 'file',
			extension: 'md',
			languageId: 'markdown'
		},
		{
			name: '.gitignore',
			type: 'file',
			extension: 'gitignore',
			languageId: 'gitignore'
		}
	];

	let expandedFolders = $state(new Set(['src'])); 

	function handleToggle(event) {
		const { item, expanded } = event.detail;
		const path = item.path; // Use the path that was already calculated
		
		console.log('Toggle:', path, 'expanded:', expanded);
		
		if (expanded) {
			expandedFolders.add(path);
		} else {
			// Remove this folder and all its descendants from expanded set
			const pathsToRemove = Array.from(expandedFolders).filter(expandedPath => 
				expandedPath === path || expandedPath.startsWith(path + '/')
			);
			pathsToRemove.forEach(pathToRemove => expandedFolders.delete(pathToRemove));
		}
		// Trigger reactivity by reassigning
		expandedFolders = new Set(expandedFolders);
		console.log('Expanded folders:', Array.from(expandedFolders));
	}

	function handleSelect(event) {
		console.log('Selected:', event.detail.item);
	}

	function getItemPath(item, parentPath = '') {
		return parentPath ? `${parentPath}/${item.name}` : item.name;
	}

	function flattenTree(items, level = 0, parentPath = '') {
		let result = [];
		
		for (const item of items) {
			const currentPath = getItemPath(item, parentPath);
			const isExpanded = expandedFolders.has(currentPath);
			
			result.push({
				...item,
				level,
				expanded: isExpanded,
				path: currentPath
			});
			
			if (item.type === 'folder' && item.children && isExpanded) {
				result.push(...flattenTree(item.children, level + 1, currentPath));
			}
		}
		
		return result;
	}

	// Use derived rune for flattened tree
	const flattenedTree = $derived(flattenTree(fileTree));
</script>

<div class="preset-glass-neutral border-primary-300 h-screen w-[20%] border-r">
	<span class="border-primary-300 block border-b p-[5.6px] font-mono text-[14px] text-white">Explorer</span>
	<div class="overflow-y-auto h-full">
		{#each flattenedTree as treeItem}
			<FileTreeItem 
				item={treeItem} 
				level={treeItem.level} 
				expanded={treeItem.expanded}
				on:toggle={handleToggle}
				on:select={handleSelect} 
			/>
		{/each}
	</div>
</div>