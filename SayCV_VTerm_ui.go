// This file was created by ui2walk and may be regenerated.
// DO NOT EDIT OR YOUR MODIFICATIONS WILL BE LOST!

package main
		
import (
	"github.com/lxn/walk"
)
		
type myDialogUI struct {
	te		*walk.TextEdit
	tw		*walk.TabWidget
	displayTabPage		*walk.TabPage
	portTabPage			*walk.TabPage
	tp3		*walk.TabPage
	tp4		*walk.TabPage
	tp5		*walk.TabPage
	tp6		*walk.TabPage
	tp7		*walk.TabPage
	tp8		*walk.TabPage
	tp9		*walk.TabPage
	tp10	*walk.TabPage
	displayAsGroupBox		*walk.GroupBox
	gb2		*walk.GroupBox
	gb3		*walk.GroupBox
	gb4		*walk.GroupBox
	gb5		*walk.GroupBox
	gb6		*walk.GroupBox
	gb7		*walk.GroupBox
	
	baudLabel			*walk.Label
	portLabel			*walk.Label
	baudCbb				*walk.ComboBox
	portCbb				*walk.ComboBox
	portChangesCb		*walk.CheckBox
	
	openPortBtn		*walk.PushButton
	splitter_1	*walk.Splitter
	splitter_2	*walk.Splitter
	
	widget		*walk.Composite
}

func (w *MyDialog) init(owner walk.Form) (err error) {
	if w.Dialog, err = walk.NewDialog(owner); err != nil {
		return err
	}

	succeeded := false
	defer func(){
		if !succeeded {
			w.Dispose()
		}
	}()
			
	var font *walk.Font
	if font == nil {
		font = nil
	}

	w.SetName("Dialog")
	
	if err := w.SetClientSize(walk.Size{800, 600}); err != nil {
		return err
	}
	
	if err := w.SetTitle(`Dialog`); err != nil {
		return err
	}

	// textEdit
	if w.ui.te, err = walk.NewTextEdit(w); err != nil {
		return err
	}
	
	w.ui.te.SetName("textEdit")
	if err := w.ui.te.SetBounds(walk.Rectangle{0, 0, 800, 300}); err != nil {
		return err
	}
			
	// tabWidget
	if w.ui.tw, err = walk.NewTabWidget(w); err != nil {
		return err
	}
	w.ui.tw.SetName("tabWidget")
	if err := w.ui.tw.SetBounds(walk.Rectangle{0, 300, 571, 268}); err != nil {
		return err
	}
			
	// tab display
	if w.ui.displayTabPage, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.displayTabPage.SetName("displayTabPage")
	if err := w.ui.displayTabPage.SetTitle(`Display`); err != nil {
		return err
	}
	if err := w.ui.tw.Pages().Add(w.ui.displayTabPage); err != nil {
		return err
	}
			
	// groupBox Display As
	if w.ui.displayAsGroupBox, err = walk.NewGroupBox(w.ui.displayTabPage); err != nil {
		return err
	}
	w.ui.displayAsGroupBox.SetName("displayAsGroupBox")
	if err := w.ui.displayAsGroupBox.SetBounds(walk.Rectangle{0, 0, 130, 241}); err != nil {
		return err
	}
	if err := w.ui.displayAsGroupBox.SetTitle(`Display As`); err != nil {
		return err
	}
	
	// comboBox
	if w.ui.baudCbb, err = walk.NewComboBox(w.ui.gb2); err != nil {
		return err
	}
	w.ui.baudCbb.SetName("comboBox")
	
	// tab_2
	if w.ui.portTabPage, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.portTabPage.SetName("tab_2")
	if err := w.ui.portTabPage.SetTitle(`Port`); err != nil {
		return err
	}
	if err := w.ui.tw.Pages().Add(w.ui.portTabPage); err != nil {
		return err
	}
			
	// groupBox_2
	if w.ui.gb2, err = walk.NewGroupBox(w.ui.portTabPage); err != nil {
		return err
	}
	w.ui.gb2.SetName("groupBox_2")
	if err := w.ui.gb2.SetBounds(walk.Rectangle{0, 0, 561, 221}); err != nil {
		return err
	}
	if err := w.ui.gb2.SetTitle(`GroupBox`); err != nil {
		return err
	}

	// label
	if w.ui.baudLabel, err = walk.NewLabel(w.ui.gb2); err != nil {
		return err
	}
	w.ui.baudLabel.SetName("label")
	if err := w.ui.baudLabel.SetBounds(walk.Rectangle{10, 30, 24, 12}); err != nil {
		return err
	}
	if err := w.ui.baudLabel.SetText(`Baud`); err != nil {
		return err
	}

	// comboBox
	if w.ui.baudCbb, err = walk.NewComboBox(w.ui.gb2); err != nil {
		return err
	}
	w.ui.baudCbb.SetName("comboBox")
	if err := w.ui.baudCbb.SetBounds(walk.Rectangle{40, 20, 74, 22}); err != nil {
		return err
	}
			
	// comboBox_2
	if w.ui.portCbb, err = walk.NewComboBox(w.ui.gb2); err != nil {
		return err
	}
	w.ui.portCbb.SetName("comboBox_2")
	if err := w.ui.portCbb.SetBounds(walk.Rectangle{160, 20, 74, 22}); err != nil {
		return err
	}
			
	// label_2
	if w.ui.portLabel, err = walk.NewLabel(w.ui.gb2); err != nil {
		return err
	}
	w.ui.portLabel.SetName("label_2")
	if err := w.ui.portLabel.SetBounds(walk.Rectangle{130, 30, 24, 12}); err != nil {
		return err
	}
	if err := w.ui.portLabel.SetText(`Port`); err != nil {
		return err
	}

	// pushButton
	if w.ui.openPortBtn, err = walk.NewPushButton(w.ui.gb2); err != nil {
		return err
	}
	w.ui.openPortBtn.SetName("pushButton")
	if err := w.ui.openPortBtn.SetBounds(walk.Rectangle{250, 20, 75, 23}); err != nil {
		return err
	}
	if err := w.ui.openPortBtn.SetText(`Open`); err != nil {
		return err
	}

	// checkBox
	if w.ui.portChangesCb, err = walk.NewCheckBox(w.ui.gb2); err != nil {
		return err
	}
	w.ui.portChangesCb.SetName("checkBox")
	if err := w.ui.portChangesCb.SetBounds(walk.Rectangle{340, 20, 13, 13}); err != nil {
		return err
	}
	if err := w.ui.portChangesCb.SetText(``); err != nil {
		return err
	}

	// splitter_1
	if w.ui.splitter_1, err = walk.NewHSplitter(w.ui.gb2); err != nil {
		return err
	}
	w.ui.splitter_1.SetName("splitter")
	if err := w.ui.splitter_1.SetBounds(walk.Rectangle{10, 70, 156, 143}); err != nil {
		return err
	}
			
	// groupBox_3
	if w.ui.gb3, err = walk.NewGroupBox(w.ui.splitter_1); err != nil {
		return err
	}
	w.ui.gb3.SetName("groupBox_3")
	if err := w.ui.gb3.SetTitle(`Parity`); err != nil {
		return err
	}

	// groupBox_4
	if w.ui.gb4, err = walk.NewGroupBox(w.ui.splitter_1); err != nil {
		return err
	}
	w.ui.gb4.SetName("groupBox_4")
	if err := w.ui.gb4.SetTitle(`Data Bits`); err != nil {
		return err
	}

	// splitter_2
	if w.ui.splitter_2, err = walk.NewHSplitter(w.ui.gb2); err != nil {
		return err
	}
	w.ui.splitter_2.SetName("splitter_2")
	if err := w.ui.splitter_2.SetBounds(walk.Rectangle{170, 70, 170, 137}); err != nil {
		return err
	}
			
	// groupBox_5
	if w.ui.gb5, err = walk.NewGroupBox(w.ui.splitter_2); err != nil {
		return err
	}
	w.ui.gb5.SetName("groupBox_5")
	if err := w.ui.gb5.SetTitle(`Stop Bits`); err != nil {
		return err
	}

	// groupBox_6
	if w.ui.gb6, err = walk.NewGroupBox(w.ui.splitter_2); err != nil {
		return err
	}
	w.ui.gb6.SetName("groupBox_6")
	if err := w.ui.gb6.SetTitle(`Hardware Flow Control`); err != nil {
		return err
	}

	// tab_3
	if w.ui.tp3, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tp3.SetName("tab_3")
	if err := w.ui.tp3.SetTitle(`Capture`); err != nil {
		return err
	}
	
	if err := w.ui.tw.Pages().Add(w.ui.tp3); err != nil {
		return err
	}
			
	// tab_4
	if w.ui.tp4, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tp4.SetName("tab_4")
	if err := w.ui.tp4.SetTitle(`Pins`); err != nil {
		return err
	}
	if err := w.ui.tw.Pages().Add(w.ui.tp4); err != nil {
		return err
	}
			
	// tab_5
	if w.ui.tp5, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tp5.SetName("tab_5")
	if err := w.ui.tp5.SetTitle(`Send`); err != nil {
		return err
	}
	if err := w.ui.tw.Pages().Add(w.ui.tp5); err != nil {
		return err
	}
			
	// tab_6
	if w.ui.tp6, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tp6.SetName("tab_6")
	if err := w.ui.tp6.SetTitle(`Echo Port`); err != nil {
		return err
	}
	if err := w.ui.tw.Pages().Add(w.ui.tp6); err != nil {
		return err
	}
			
	// tab_7
	if w.ui.tp7, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tp7.SetName("tab_7")
	if err := w.ui.tp7.SetTitle(`I2C`); err != nil {
		return err
	}
	if err := w.ui.tw.Pages().Add(w.ui.tp7); err != nil {
		return err
	}
			
	// tab_8
	if w.ui.tp8, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tp8.SetName("tab_8")
	if err := w.ui.tp8.SetTitle(`I2C-2`); err != nil {
		return err
	}
	if err := w.ui.tw.Pages().Add(w.ui.tp8); err != nil {
		return err
	}
			
	// tab_9
	if w.ui.tp9, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tp9.SetName("tab_9")
	if err := w.ui.tp9.SetTitle(`I2CMisc`); err != nil {
		return err
	}
	if err := w.ui.tw.Pages().Add(w.ui.tp9); err != nil {
		return err
	}
			
	// tab_10
	if w.ui.tp10, err = walk.NewTabPage(); err != nil {
		return err
	}
	w.ui.tp10.SetName("tab_10")
	if err := w.ui.tp10.SetTitle(`Misc`); err != nil {
		return err
	}
	if err := w.ui.tw.Pages().Add(w.ui.tp10); err != nil {
		return err
	}
			
	// widget
	if w.ui.widget, err = walk.NewComposite(w); err != nil {
		return err
	}
	w.ui.widget.SetName("widget")
	if err := w.ui.widget.SetBounds(walk.Rectangle{0, 568, 800, 32}); err != nil {
		return err
	}
			
	// groupBox_7
	if w.ui.gb7, err = walk.NewGroupBox(w); err != nil {
		return err
	}
	w.ui.gb7.SetName("groupBox_7")
	if err := w.ui.gb7.SetBounds(walk.Rectangle{580, 360, 120, 211}); err != nil {
		return err
	}
	if err := w.ui.gb7.SetTitle(`Status`); err != nil {
		return err
	}
	
	// Tab order

	succeeded = true
		
	return nil
}
