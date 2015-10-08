_         = require 'lodash'
kd        = require 'kd'
JView     = require 'app/jview'
remote    = require('app/remote').getInstance()
showError = require 'app/util/showError'
Validator = require 'validator'


module.exports = class TeamInviteView extends JView

  constructor:(options, data)->

    super options, data

    @textarea = new kd.InputView
      type        : 'textarea'
      placeholder : """
        Put email addresses comma/newline separated e.g.
        sinan@koding.com, devrim@koding.com, nitin@koding.com

        or:

        sinan@koding.com
        devrim@koding.com

        or go crazy and try me:

        senthil@koding.com, indian-rider@software-engineer-girl.com
        --- oh@yeah I can break things .com ---
        nitin@koding.com, nicolo@koding.com
        , some jibberish @@#!
        """

    @confirmationView = new kd.ScrollView

    @invite = new kd.ButtonView
      title    : 'Verify emails'
      cssClass : 'solid medium green fr'
      callback : @bound 'prepareEmails'

    @confirm = new kd.ButtonView
      title    : 'Looks good, Send\'em ALL!'
      cssClass : 'solid medium green fr hidden'
      callback : @bound 'sendInvites'

    @abort = new kd.ButtonView
      title    : 'Abort'
      cssClass : 'solid medium red fr hidden'
      callback : @bound 'hideConfirmation'


  prepareEmails: ->

    rawText  = @textarea.getValue()
    rawText  = rawText.replace /\n/g, ','
    rawText  = rawText.replace /\s/g, ''
    splitted = rawText.split ','
    splitted = _.uniq splitted
    @emails  = _.remove splitted, (email) -> Validator.isEmail email

    @showConfirmation()


  sendInvites: ->

    remote.api.JTeamInvitation.sendInvitationEmails @emails, =>
      new kd.NotificationView title : 'Invitations sent!'
      @hideConfirmation()


  showConfirmation: ->

    return new kd.NotificationView title : 'You\'re doing it wrong!'  unless @emails.length > 0

    @confirmationView.updatePartial @emails.join '<br/>'

    @textarea.hide()
    @invite.hide()

    @confirmationView.show()
    @confirm.show()
    @abort.show()


  hideConfirmation: ->

    @confirmationView.updatePartial ''

    @confirmationView.hide()
    @confirm.hide()
    @abort.hide()

    @textarea.show()
    @invite.show()



  pistachio: ->

    """
    {{> @textarea}}
    {{> @invite}}
    {{> @confirmationView}}
    {{> @confirm}}{{> @abort}}
    """