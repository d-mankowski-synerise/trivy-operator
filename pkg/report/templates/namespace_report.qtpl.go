// Code generated by qtc from "namespace_report.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pkg/report/templates/namespace_report.qtpl:1
package templates

//line pkg/report/templates/namespace_report.qtpl:1
import "github.com/aquasecurity/trivy-operator/pkg/apis/aquasecurity/v1alpha1"

//line pkg/report/templates/namespace_report.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pkg/report/templates/namespace_report.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pkg/report/templates/namespace_report.qtpl:3
func (p *NamespaceReport) StreamTitle(qw422016 *qt422016.Writer) {
//line pkg/report/templates/namespace_report.qtpl:3
	qw422016.N().S(`
Aqua Starboard Namespace Security Report - Namespace: `)
//line pkg/report/templates/namespace_report.qtpl:4
	qw422016.E().S(p.Namespace.Name)
//line pkg/report/templates/namespace_report.qtpl:4
	qw422016.N().S(`
`)
//line pkg/report/templates/namespace_report.qtpl:5
}

//line pkg/report/templates/namespace_report.qtpl:5
func (p *NamespaceReport) WriteTitle(qq422016 qtio422016.Writer) {
//line pkg/report/templates/namespace_report.qtpl:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/report/templates/namespace_report.qtpl:5
	p.StreamTitle(qw422016)
//line pkg/report/templates/namespace_report.qtpl:5
	qt422016.ReleaseWriter(qw422016)
//line pkg/report/templates/namespace_report.qtpl:5
}

//line pkg/report/templates/namespace_report.qtpl:5
func (p *NamespaceReport) Title() string {
//line pkg/report/templates/namespace_report.qtpl:5
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/report/templates/namespace_report.qtpl:5
	p.WriteTitle(qb422016)
//line pkg/report/templates/namespace_report.qtpl:5
	qs422016 := string(qb422016.B)
//line pkg/report/templates/namespace_report.qtpl:5
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/report/templates/namespace_report.qtpl:5
	return qs422016
//line pkg/report/templates/namespace_report.qtpl:5
}

//line pkg/report/templates/namespace_report.qtpl:7
func (p *NamespaceReport) StreamBody(qw422016 *qt422016.Writer) {
//line pkg/report/templates/namespace_report.qtpl:7
	qw422016.N().S(`
<div class="container">

  <div class="col mt-5">
    <div class="row text-center">`)
//line pkg/report/templates/namespace_report.qtpl:11
	streamimgAquaLogo(qw422016)
//line pkg/report/templates/namespace_report.qtpl:11
	qw422016.N().S(`</div>
    <div class="row mt-4 text-center">
      <h2 class="text-muted mx-auto">Aqua Starboard Namespace Security Report</h2>
    </div>
    <div class="row text-center">
      <h3 class="text-muted mx-auto">`)
//line pkg/report/templates/namespace_report.qtpl:16
	qw422016.E().S(string(p.Namespace.Kind))
//line pkg/report/templates/namespace_report.qtpl:16
	qw422016.N().S(`: `)
//line pkg/report/templates/namespace_report.qtpl:16
	qw422016.E().S(p.Namespace.Name)
//line pkg/report/templates/namespace_report.qtpl:16
	qw422016.N().S(`</h3>
    </div>
    <div class="row text-center">
      <h3 class="text-muted mx-auto">Generated on `)
//line pkg/report/templates/namespace_report.qtpl:19
	qw422016.E().S(p.GeneratedAt.Format("2 Jan 2006 15:04:01"))
//line pkg/report/templates/namespace_report.qtpl:19
	qw422016.N().S(`</h3>
    </div>
  </div>

  <div class="row">
    <h3>Top 5 vulnerable images by count</h3>
    <table class="table table-sm table-bordered">
      <thead>
        <tr>
          <th scope="col">Image</th>
          <th scope="col">Critical</th>
          <th scope="col">High</th>
          <th scope="col">Other</th>
        </tr>
      </thead>
      <tbody>
      `)
//line pkg/report/templates/namespace_report.qtpl:35
	for _, report := range p.Top5VulnerableImages {
//line pkg/report/templates/namespace_report.qtpl:35
		qw422016.N().S(`
      `)
//line pkg/report/templates/namespace_report.qtpl:37
		summary := report.Report.Summary
		otherCount := summary.MediumCount + summary.LowCount + summary.UnknownCount

//line pkg/report/templates/namespace_report.qtpl:39
		qw422016.N().S(`
      <tr>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:41
		streamimageReference(qw422016, report.Report.Registry, report.Report.Artifact)
//line pkg/report/templates/namespace_report.qtpl:41
		qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:42
		qw422016.N().D(summary.CriticalCount)
//line pkg/report/templates/namespace_report.qtpl:42
		qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:43
		qw422016.N().D(summary.HighCount)
//line pkg/report/templates/namespace_report.qtpl:43
		qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:44
		qw422016.N().D(otherCount)
//line pkg/report/templates/namespace_report.qtpl:44
		qw422016.N().S(`</td>
      </tr>
      `)
//line pkg/report/templates/namespace_report.qtpl:46
	}
//line pkg/report/templates/namespace_report.qtpl:46
	qw422016.N().S(`
      </tbody>
    </table>
  </div>

  <div class="row">
    <h3>Top 5 vulnerabilities by score</h3>
    <table class="table table-sm table-bordered">
      <thead>
        <tr>
          <th scope="col">ID</th>
          <th scope="col">Severity</th>
          <th scope="col">Score</th>
          <th scope="col">Affected Workloads</th>
        </tr>
      </thead>
      <tbody>
      `)
//line pkg/report/templates/namespace_report.qtpl:63
	for _, vulnerability := range p.Top5Vulnerability {
//line pkg/report/templates/namespace_report.qtpl:63
		qw422016.N().S(`
      <tr>
        <td><a href="`)
//line pkg/report/templates/namespace_report.qtpl:65
		qw422016.E().S(vulnerability.PrimaryLink)
//line pkg/report/templates/namespace_report.qtpl:65
		qw422016.N().S(`">`)
//line pkg/report/templates/namespace_report.qtpl:65
		qw422016.E().S(vulnerability.VulnerabilityID)
//line pkg/report/templates/namespace_report.qtpl:65
		qw422016.N().S(`</a></td>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:66
		qw422016.E().S(string(vulnerability.Severity))
//line pkg/report/templates/namespace_report.qtpl:66
		qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:67
		qw422016.N().F(*vulnerability.Score)
//line pkg/report/templates/namespace_report.qtpl:67
		qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:68
		qw422016.N().D(vulnerability.AffectedWorkloads)
//line pkg/report/templates/namespace_report.qtpl:68
		qw422016.N().S(`</td>
      </tr>
      `)
//line pkg/report/templates/namespace_report.qtpl:70
	}
//line pkg/report/templates/namespace_report.qtpl:70
	qw422016.N().S(`
      </tbody>
    </table>
  </div>

  <div class="row">
    <h3>Top 5 failed workload configs</h3>
    <table class="table table-sm table-bordered">
      <thead>
        <tr>
          <th scope="col">Name</th>
          <th scope="col">Severity</th>
          <th scope="col">Category</th>
          <th scope="col">Affected Workloads</th>
        </tr>
      </thead>
      <tbody>
      `)
//line pkg/report/templates/namespace_report.qtpl:87
	for _, report := range p.Top5FailedChecks {
//line pkg/report/templates/namespace_report.qtpl:87
		qw422016.N().S(`
      <tr>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:89
		qw422016.E().S(report.ID)
//line pkg/report/templates/namespace_report.qtpl:89
		qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:90
		qw422016.E().V(report.Severity)
//line pkg/report/templates/namespace_report.qtpl:90
		qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:91
		qw422016.E().S(report.Category)
//line pkg/report/templates/namespace_report.qtpl:91
		qw422016.N().S(`</td>
        <td>`)
//line pkg/report/templates/namespace_report.qtpl:92
		qw422016.N().D(report.AffectedWorkloads)
//line pkg/report/templates/namespace_report.qtpl:92
		qw422016.N().S(`</td>
      </tr>
      `)
//line pkg/report/templates/namespace_report.qtpl:94
	}
//line pkg/report/templates/namespace_report.qtpl:94
	qw422016.N().S(`
      </tbody>
    </table>
  </div>

</div>
`)
//line pkg/report/templates/namespace_report.qtpl:100
}

//line pkg/report/templates/namespace_report.qtpl:100
func (p *NamespaceReport) WriteBody(qq422016 qtio422016.Writer) {
//line pkg/report/templates/namespace_report.qtpl:100
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/report/templates/namespace_report.qtpl:100
	p.StreamBody(qw422016)
//line pkg/report/templates/namespace_report.qtpl:100
	qt422016.ReleaseWriter(qw422016)
//line pkg/report/templates/namespace_report.qtpl:100
}

//line pkg/report/templates/namespace_report.qtpl:100
func (p *NamespaceReport) Body() string {
//line pkg/report/templates/namespace_report.qtpl:100
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/report/templates/namespace_report.qtpl:100
	p.WriteBody(qb422016)
//line pkg/report/templates/namespace_report.qtpl:100
	qs422016 := string(qb422016.B)
//line pkg/report/templates/namespace_report.qtpl:100
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/report/templates/namespace_report.qtpl:100
	return qs422016
//line pkg/report/templates/namespace_report.qtpl:100
}

//line pkg/report/templates/namespace_report.qtpl:102
func streamimageReference(qw422016 *qt422016.Writer, registry v1alpha1.Registry, artifact v1alpha1.Artifact) {
//line pkg/report/templates/namespace_report.qtpl:102
	qw422016.N().S(`
  `)
//line pkg/report/templates/namespace_report.qtpl:103
	if artifact.Tag != "" && artifact.Digest != "" {
//line pkg/report/templates/namespace_report.qtpl:103
		qw422016.N().S(`
    `)
//line pkg/report/templates/namespace_report.qtpl:104
		qw422016.E().S(registry.Server)
//line pkg/report/templates/namespace_report.qtpl:104
		qw422016.N().S(`/`)
//line pkg/report/templates/namespace_report.qtpl:104
		qw422016.E().S(artifact.Repository)
//line pkg/report/templates/namespace_report.qtpl:104
		qw422016.N().S(`:`)
//line pkg/report/templates/namespace_report.qtpl:104
		qw422016.E().S(artifact.Tag)
//line pkg/report/templates/namespace_report.qtpl:104
		qw422016.N().S(`@`)
//line pkg/report/templates/namespace_report.qtpl:104
		qw422016.E().S(artifact.Digest)
//line pkg/report/templates/namespace_report.qtpl:104
		qw422016.N().S(`
    `)
//line pkg/report/templates/namespace_report.qtpl:105
		return
//line pkg/report/templates/namespace_report.qtpl:106
	}
//line pkg/report/templates/namespace_report.qtpl:106
	qw422016.N().S(`

  `)
//line pkg/report/templates/namespace_report.qtpl:108
	if artifact.Tag == "" && artifact.Digest != "" {
//line pkg/report/templates/namespace_report.qtpl:108
		qw422016.N().S(`
    `)
//line pkg/report/templates/namespace_report.qtpl:109
		qw422016.E().S(registry.Server)
//line pkg/report/templates/namespace_report.qtpl:109
		qw422016.N().S(`/`)
//line pkg/report/templates/namespace_report.qtpl:109
		qw422016.E().S(artifact.Repository)
//line pkg/report/templates/namespace_report.qtpl:109
		qw422016.N().S(`@`)
//line pkg/report/templates/namespace_report.qtpl:109
		qw422016.E().S(artifact.Digest)
//line pkg/report/templates/namespace_report.qtpl:109
		qw422016.N().S(`
    `)
//line pkg/report/templates/namespace_report.qtpl:110
		return
//line pkg/report/templates/namespace_report.qtpl:111
	}
//line pkg/report/templates/namespace_report.qtpl:111
	qw422016.N().S(`

  `)
//line pkg/report/templates/namespace_report.qtpl:113
	if artifact.Tag != "" && artifact.Digest == "" {
//line pkg/report/templates/namespace_report.qtpl:113
		qw422016.N().S(`
    `)
//line pkg/report/templates/namespace_report.qtpl:114
		qw422016.E().S(registry.Server)
//line pkg/report/templates/namespace_report.qtpl:114
		qw422016.N().S(`/`)
//line pkg/report/templates/namespace_report.qtpl:114
		qw422016.E().S(artifact.Repository)
//line pkg/report/templates/namespace_report.qtpl:114
		qw422016.N().S(`:`)
//line pkg/report/templates/namespace_report.qtpl:114
		qw422016.E().S(artifact.Tag)
//line pkg/report/templates/namespace_report.qtpl:114
		qw422016.N().S(`
    `)
//line pkg/report/templates/namespace_report.qtpl:115
		return
//line pkg/report/templates/namespace_report.qtpl:116
	}
//line pkg/report/templates/namespace_report.qtpl:116
	qw422016.N().S(`

  `)
//line pkg/report/templates/namespace_report.qtpl:118
	qw422016.E().S(registry.Server)
//line pkg/report/templates/namespace_report.qtpl:118
	qw422016.N().S(`/`)
//line pkg/report/templates/namespace_report.qtpl:118
	qw422016.E().S(artifact.Repository)
//line pkg/report/templates/namespace_report.qtpl:118
	qw422016.N().S(`:`)
//line pkg/report/templates/namespace_report.qtpl:118
	qw422016.E().S(artifact.Tag)
//line pkg/report/templates/namespace_report.qtpl:118
	qw422016.N().S(`
`)
//line pkg/report/templates/namespace_report.qtpl:119
}

//line pkg/report/templates/namespace_report.qtpl:119
func writeimageReference(qq422016 qtio422016.Writer, registry v1alpha1.Registry, artifact v1alpha1.Artifact) {
//line pkg/report/templates/namespace_report.qtpl:119
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/report/templates/namespace_report.qtpl:119
	streamimageReference(qw422016, registry, artifact)
//line pkg/report/templates/namespace_report.qtpl:119
	qt422016.ReleaseWriter(qw422016)
//line pkg/report/templates/namespace_report.qtpl:119
}

//line pkg/report/templates/namespace_report.qtpl:119
func imageReference(registry v1alpha1.Registry, artifact v1alpha1.Artifact) string {
//line pkg/report/templates/namespace_report.qtpl:119
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/report/templates/namespace_report.qtpl:119
	writeimageReference(qb422016, registry, artifact)
//line pkg/report/templates/namespace_report.qtpl:119
	qs422016 := string(qb422016.B)
//line pkg/report/templates/namespace_report.qtpl:119
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/report/templates/namespace_report.qtpl:119
	return qs422016
//line pkg/report/templates/namespace_report.qtpl:119
}
