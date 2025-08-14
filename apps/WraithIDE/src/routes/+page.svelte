<script>
	import Editor from '../components/Editor.svelte';
	import Terminal from '../components/Terminal.svelte';
	import FileTree from '../components/FileTree.svelte';
</script>

<div class="flex h-screen w-screen">
	<FileTree />
	<div class="flex flex-col w-[80%] h-full">
		<Editor />
		<Terminal />
	</div>
</div>
